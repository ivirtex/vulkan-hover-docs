# VK\_KHR\_separate\_depth\_stencil\_layouts(3) Manual Page

## Name

VK\_KHR\_separate\_depth\_stencil\_layouts - device extension



## [](#_registered_extension_number)Registered Extension Number

242

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

         [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
         or  
         [Vulkan Version 1.1](#versions-1.1)  
     and  
     [VK\_KHR\_create\_renderpass2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_create_renderpass2.html)  
or  
[Vulkan Version 1.2](#versions-1.2)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_separate_depth_stencil_layouts%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_separate_depth_stencil_layouts%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-06-25

**Contributors**

- Daniel Koch, NVIDIA
- Jeff Bolz, NVIDIA
- Jesse Barker, Unity
- Tobias Hector, AMD

## [](#_description)Description

This extension allows image memory barriers for depth/stencil images to have just one of the `VK_IMAGE_ASPECT_DEPTH_BIT` or `VK_IMAGE_ASPECT_STENCIL_BIT` aspect bits set, rather than require both. This allows their layouts to be set independently. To support depth/stencil images with different layouts for the depth and stencil aspects, the depth/stencil attachment interface has been updated to support a separate layout for stencil.

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_structures)New Structures

- Extending [VkAttachmentDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescription2.html):
  
  - [VkAttachmentDescriptionStencilLayoutKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescriptionStencilLayoutKHR.html)
- Extending [VkAttachmentReference2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReference2.html):
  
  - [VkAttachmentReferenceStencilLayoutKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentReferenceStencilLayoutKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSeparateDepthStencilLayoutsFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SEPARATE_DEPTH_STENCIL_LAYOUTS_EXTENSION_NAME`
- `VK_KHR_SEPARATE_DEPTH_STENCIL_LAYOUTS_SPEC_VERSION`
- Extending [VkImageLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageLayout.html):
  
  - `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_OPTIMAL_KHR`
  - `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_OPTIMAL_KHR`
  - `VK_IMAGE_LAYOUT_STENCIL_ATTACHMENT_OPTIMAL_KHR`
  - `VK_IMAGE_LAYOUT_STENCIL_READ_ONLY_OPTIMAL_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_ATTACHMENT_DESCRIPTION_STENCIL_LAYOUT_KHR`
  - `VK_STRUCTURE_TYPE_ATTACHMENT_REFERENCE_STENCIL_LAYOUT_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SEPARATE_DEPTH_STENCIL_LAYOUTS_FEATURES_KHR`

## [](#_version_history)Version History

- Revision 1, 2019-06-25 (Piers Daniell)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_separate_depth_stencil_layouts).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0