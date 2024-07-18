# VK_QCOM_rotated_copy_commands(3) Manual Page

## Name

VK_QCOM_rotated_copy_commands - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

334

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_copy_commands2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_copy_commands2.html)  
or  
[Version 1.3](#versions-1.3)  

## <a href="#_contact" class="anchor"></a>Contact

- Matthew Netsch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_rotated_copy_commands%5D%20@mnetsch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_rotated_copy_commands%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mnetsch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-12-13

**Interactions and External Dependencies**  
- This extension interacts with
  [`VK_KHR_swapchain`](VK_KHR_swapchain.html)

- This extension interacts with [`VK_KHR_surface`](VK_KHR_surface.html)

**Contributors**  
- Jeff Leger, Qualcomm Technologies, Inc.

- Matthew Netsch, Qualcomm Technologies, Inc.

## <a href="#_description" class="anchor"></a>Description

This extension extends adds an optional rotation transform to copy
commands [vkCmdBlitImage2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBlitImage2KHR.html),
[vkCmdCopyImageToBuffer2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyImageToBuffer2KHR.html) and
[vkCmdCopyBufferToImage2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyBufferToImage2KHR.html). When
copying between two resources, where one resource contains rotated
content and the other does not, a rotated copy may be desired. This
extension may be used in combination with VK_QCOM_render_pass_transform
which adds rotated render passes.

This extension adds an extension structure to the following commands:
vkCmdBlitImage2KHR, vkCmdCopyImageToBuffer2KHR and
vkCmdCopyBufferToImage2KHR

## <a href="#_issues" class="anchor"></a>Issues

1\) What is an appropriate name for the added extension structure? The
style guide says “Structures which extend other structures through the
`pNext` chain should reflect the name of the base structure they
extend.”, but in this case a single extension structure is used to
extend three base structures (vkCmdBlitImage2KHR,
vkCmdCopyImageToBuffer2KHR and vkCmdCopyBufferToImage2KHR). Creating
three identical structures with unique names seemed undesirable.

**RESOLVED**: Deviate from the style guide for extension structure
naming.

2\) Should this extension add a rotation capability to
vkCmdCopyImage2KHR?

**RESOLVED**: No. Use of rotated vkCmdBlitImage2KHR can fully address
this use case.

3\) Should this extension add a rotation capability to
vkCmdResolveImage2KHR?

**RESOLVED** No. Use of vkCmdResolveImage2KHR is very slow and extremely
bandwidth intensive on Qualcomm’s GPU architecture and use of
pResolveAttachments in vkRenderPass is the strongly preferred approach.
Therefore, we choose not to introduce a rotation capability to
vkCmdResolveImage2KHR.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkBufferImageCopy2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferImageCopy2.html),
  [VkImageBlit2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageBlit2.html):

  - [VkCopyCommandTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyCommandTransformInfoQCOM.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_QCOM_ROTATED_COPY_COMMANDS_EXTENSION_NAME`

- `VK_QCOM_ROTATED_COPY_COMMANDS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_COPY_COMMAND_TRANSFORM_INFO_QCOM`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-09-19 (Jeff Leger)

- Revision 2, 2023-12-13 (Matthew Netsch)

  - Relax dependency on VK_KHR_swapchain

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_QCOM_rotated_copy_commands"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
