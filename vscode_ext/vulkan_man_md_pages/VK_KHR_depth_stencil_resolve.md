# VK_KHR_depth_stencil_resolve(3) Manual Page

## Name

VK_KHR_depth_stencil_resolve - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

200

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_create_renderpass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_create_renderpass2.html)  
or  
[Version 1.2](#versions-1.2)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Jan-Harald Fredriksen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_depth_stencil_resolve%5D%20@janharald%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_depth_stencil_resolve%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>janharald</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-04-09

**Contributors**  
- Jan-Harald Fredriksen, Arm

- Andrew Garrard, Samsung Electronics

- Soowan Park, Samsung Electronics

- Jeff Bolz, NVIDIA

- Daniel Rakos, AMD

## <a href="#_description" class="anchor"></a>Description

This extension adds support for automatically resolving multisampled
depth/stencil attachments in a subpass in a similar manner as for color
attachments.

Multisampled color attachments can be resolved at the end of a subpass
by specifying `pResolveAttachments` entries corresponding to the
`pColorAttachments` array entries. This does not allow for a way to map
the resolve attachments to the depth/stencil attachment. The
[vkCmdResolveImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdResolveImage.html) command does not allow for
depth/stencil images. While there are other ways to resolve the
depth/stencil attachment, they can give sub-optimal performance.
Extending the `VkSubpassDescription2` in this extension allows an
application to add a `pDepthStencilResolveAttachment`, that is similar
to the color `pResolveAttachments`, that the `pDepthStencilAttachment`
can be resolved into.

Depth and stencil samples are resolved to a single value based on the
resolve mode. The set of possible resolve modes is defined in the
[VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html) enum. The
`VK_RESOLVE_MODE_SAMPLE_ZERO_BIT` mode is the only mode that is required
of all implementations (that support the extension or support Vulkan 1.2
or higher). Some implementations may also support averaging (the same as
color sample resolve) or taking the minimum or maximum sample, which may
be more suitable for depth/stencil resolve.

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceDepthStencilResolvePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolvePropertiesKHR.html)

- Extending [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html):

  - [VkSubpassDescriptionDepthStencilResolveKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionDepthStencilResolveKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkResolveModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBitsKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkResolveModeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_DEPTH_STENCIL_RESOLVE_EXTENSION_NAME`

- `VK_KHR_DEPTH_STENCIL_RESOLVE_SPEC_VERSION`

- Extending [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlagBits.html):

  - `VK_RESOLVE_MODE_AVERAGE_BIT_KHR`

  - `VK_RESOLVE_MODE_MAX_BIT_KHR`

  - `VK_RESOLVE_MODE_MIN_BIT_KHR`

  - `VK_RESOLVE_MODE_NONE_KHR`

  - `VK_RESOLVE_MODE_SAMPLE_ZERO_BIT_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE_KHR`

## <a href="#_additional_resources" class="anchor"></a>Additional Resources

- GDC 2019

  - [`video`](https://www.youtube.com/watch?v=GnnEmJFFC7Q&feature=youtu.be&t=1983)

  - [`slides`](https://www.khronos.org/assets/uploads/developers/library/2019-gdc/Vulkan-Depth-Stencil-Resolve-GDC-Mar19.pdf)

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-04-09 (Jan-Harald Fredriksen)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_depth_stencil_resolve"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
