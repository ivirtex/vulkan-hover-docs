# VK\_KHR\_imageless\_framebuffer(3) Manual Page

## Name

VK\_KHR\_imageless\_framebuffer - device extension



## [](#_registered_extension_number)Registered Extension Number

109

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

             [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
             and  
             [VK\_KHR\_maintenance2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance2.html)  
         or  
         [Vulkan Version 1.1](#versions-1.1)  
     and  
     [VK\_KHR\_image\_format\_list](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_image_format_list.html)  
or  
[Vulkan Version 1.2](#versions-1.2)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Tobias Hector [\[GitHub\]tobias](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_imageless_framebuffer%5D%20%40tobias%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_imageless_framebuffer%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-12-14

**Contributors**

- Tobias Hector
- Graham Wihlidal

## [](#_description)Description

This extension allows framebuffers to be created without the need for creating images first, allowing more flexibility in how they are used, and avoiding the need for many of the confusing compatibility rules.

Framebuffers are now created with a small amount of additional metadata about the image views that will be used in [VkFramebufferAttachmentsCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfoKHR.html), and the actual image views are provided at render pass begin time via [VkRenderPassAttachmentBeginInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfoKHR.html).

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_structures)New Structures

- [VkFramebufferAttachmentImageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentImageInfoKHR.html)
- Extending [VkFramebufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateInfo.html):
  
  - [VkFramebufferAttachmentsCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferAttachmentsCreateInfoKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceImagelessFramebufferFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImagelessFramebufferFeaturesKHR.html)
- Extending [VkRenderPassBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassBeginInfo.html):
  
  - [VkRenderPassAttachmentBeginInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassAttachmentBeginInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_IMAGELESS_FRAMEBUFFER_EXTENSION_NAME`
- `VK_KHR_IMAGELESS_FRAMEBUFFER_SPEC_VERSION`
- Extending [VkFramebufferCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferCreateFlagBits.html):
  
  - `VK_FRAMEBUFFER_CREATE_IMAGELESS_BIT_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENTS_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_FRAMEBUFFER_ATTACHMENT_IMAGE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_IMAGELESS_FRAMEBUFFER_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_RENDER_PASS_ATTACHMENT_BEGIN_INFO_KHR`

## [](#_version_history)Version History

- Revision 1, 2018-12-14 (Tobias Hector)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_imageless_framebuffer)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0