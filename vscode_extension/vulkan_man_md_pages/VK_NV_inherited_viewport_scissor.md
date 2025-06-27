# VK_NV_inherited_viewport_scissor(3) Manual Page

## Name

VK_NV_inherited_viewport_scissor - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

279

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- David Zhao Akeley <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_inherited_viewport_scissor%5D%20@akeley98%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_inherited_viewport_scissor%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>akeley98</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-02-04

**Contributors**  
- David Zhao Akeley, NVIDIA

- Jeff Bolz, NVIDIA

- Piers Daniell, NVIDIA

- Christoph Kubisch, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds the ability for a secondary command buffer to
inherit the dynamic viewport and scissor state from a primary command
buffer, or a previous secondary command buffer executed within the same
[vkCmdExecuteCommands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdExecuteCommands.html) call. It addresses a
frequent scenario in applications that deal with window resizing and
want to improve utilization of reusable secondary command buffers. The
functionality is provided through
[VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html).
Viewport inheritance is effectively limited to the 2D rectangle;
secondary command buffers must re-specify the inherited depth range
values.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html):

  - [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceInheritedViewportScissorFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceInheritedViewportScissorFeaturesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_INHERITED_VIEWPORT_SCISSOR_EXTENSION_NAME`

- `VK_NV_INHERITED_VIEWPORT_SCISSOR_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_VIEWPORT_SCISSOR_INFO_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INHERITED_VIEWPORT_SCISSOR_FEATURES_NV`

## <a href="#_issues" class="anchor"></a>Issues

\(1\) Why are viewport depth values configured in the
[VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html)
struct, rather than by a `vkCmd…​` function?

**DISCUSSION**:

We considered both adding a new `vkCmdSetViewportDepthNV` function, and
modifying [vkCmdSetViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewport.html) to ignore the `x`,
`y`, `width`, and `height` values when called with a secondary command
buffer that activates this extension.

The primary design considerations for this extension are debuggability
and easy integration into existing applications. The main issue with
adding a new `vkCmdSetViewportDepthNV` function is reducing
ease-of-integration. A new function pointer will have to be loaded, but
more importantly, a new function would require changes to be supported
in graphics debuggers; this would delay widespread adoption of the
extension.

The proposal to modify [vkCmdSetViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewport.html) would
avoid these issues. However, we expect that the intent of applications
using this extension is to have the viewport values used for drawing
exactly match the inherited values; thus, it would be better for
debuggability if no function for modifying the viewport depth alone is
provided. By specifying viewport depth values when starting secondary
command buffer recording, and requiring the specified depth values to
match the inherited depth values, we allow for validation layers that
flag depth changes as errors.

This design also better matches the hardware model. In fact, there is no
need to re-execute a depth-setting command. The graphics device retains
the viewport depth state; it is the CPU-side state of
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) that must be re-initialized.

\(2\) Why are viewport depth values specified as a partial
[VkViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html) struct, rather than a leaner depth-only
struct?

**DISCUSSION**:

We considered adding a new `VkViewportDepthNV` struct containing only
`minDepth` and `maxDepth`. However, as application developers would need
to maintain both a `VK_NV_inherited_viewport_scissor` code path and a
fallback code path (at least in the short term), we ultimately chose to
continue using the existing [VkViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html) structure.
Doing so would allow application developers to reuse the same
[VkViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html) array for both code paths, rather than
constructing separate `VkViewportDepthNV` and
[VkViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html) arrays for each code path.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-02-04 (David Zhao Akeley)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_inherited_viewport_scissor"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
