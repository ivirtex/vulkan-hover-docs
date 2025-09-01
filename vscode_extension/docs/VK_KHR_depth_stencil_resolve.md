# VK\_KHR\_depth\_stencil\_resolve(3) Manual Page

## Name

VK\_KHR\_depth\_stencil\_resolve - device extension



## [](#_registered_extension_number)Registered Extension Number

200

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_create\_renderpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_create_renderpass2.html)  
or  
[Vulkan Version 1.2](#versions-1.2)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Jan-Harald Fredriksen [\[GitHub\]janharald](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_depth_stencil_resolve%5D%20%40janharald%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_depth_stencil_resolve%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-04-09

**Contributors**

- Jan-Harald Fredriksen, Arm
- Andrew Garrard, Samsung Electronics
- Soowan Park, Samsung Electronics
- Jeff Bolz, NVIDIA
- Daniel Rakos, AMD

## [](#_description)Description

This extension adds support for automatically resolving multisampled depth/stencil attachments in a subpass in a similar manner as for color attachments.

Multisampled color attachments can be resolved at the end of a subpass by specifying `pResolveAttachments` entries corresponding to the `pColorAttachments` array entries. This does not allow for a way to map the resolve attachments to the depth/stencil attachment. The [vkCmdResolveImage](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdResolveImage.html) command does not allow for depth/stencil images. While there are other ways to resolve the depth/stencil attachment, they can give sub-optimal performance. Extending the `VkSubpassDescription2` in this extension allows an application to add a `pDepthStencilResolveAttachment`, that is similar to the color `pResolveAttachments`, that the `pDepthStencilAttachment` can be resolved into.

Depth and stencil samples are resolved to a single value based on the resolve mode. The set of possible resolve modes is defined in the [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlagBits.html) enum. The `VK_RESOLVE_MODE_SAMPLE_ZERO_BIT` mode is the only mode that is required of all implementations (that support the extension or support Vulkan 1.2 or higher). Some implementations may also support averaging (the same as color sample resolve) or taking the minimum or maximum sample, which may be more suitable for depth/stencil resolve.

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceDepthStencilResolvePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthStencilResolvePropertiesKHR.html)
- Extending [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html):
  
  - [VkSubpassDescriptionDepthStencilResolveKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionDepthStencilResolveKHR.html)

## [](#_new_enums)New Enums

- [VkResolveModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlagBitsKHR.html)

## [](#_new_bitmasks)New Bitmasks

- [VkResolveModeFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlagsKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_DEPTH_STENCIL_RESOLVE_EXTENSION_NAME`
- `VK_KHR_DEPTH_STENCIL_RESOLVE_SPEC_VERSION`
- Extending [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlagBits.html):
  
  - `VK_RESOLVE_MODE_AVERAGE_BIT_KHR`
  - `VK_RESOLVE_MODE_MAX_BIT_KHR`
  - `VK_RESOLVE_MODE_MIN_BIT_KHR`
  - `VK_RESOLVE_MODE_NONE_KHR`
  - `VK_RESOLVE_MODE_SAMPLE_ZERO_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_STENCIL_RESOLVE_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_SUBPASS_DESCRIPTION_DEPTH_STENCIL_RESOLVE_KHR`

## [](#_additional_resources)Additional Resources

- GDC 2019
  
  - [`video`](https://www.youtube.com/watch?v=GnnEmJFFC7Q&feature=youtu.be&t=1983)
  - [`slides`](https://www.khronos.org/assets/uploads/developers/library/2019-gdc/Vulkan-Depth-Stencil-Resolve-GDC-Mar19.pdf)

## [](#_version_history)Version History

- Revision 1, 2018-04-09 (Jan-Harald Fredriksen)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_depth_stencil_resolve).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0