# VK\_QNX\_external\_memory\_screen\_buffer(3) Manual Page

## Name

VK\_QNX\_external\_memory\_screen\_buffer - device extension



## [](#_registered_extension_number)Registered Extension Number

530

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

         [VK\_KHR\_sampler\_ycbcr\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_sampler_ycbcr_conversion.html)  
         and  
         [VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html)  
         and  
         [VK\_KHR\_dedicated\_allocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dedicated_allocation.html)  
     or  
     [Vulkan Version 1.1](#versions-1.1)  
and  
[VK\_EXT\_queue\_family\_foreign](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_queue_family_foreign.html)

## [](#_contact)Contact

- Mike Gorchak [\[GitHub\]mgorchak-blackberry](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QNX_external_memory_screen_buffer%5D%20%40mgorchak-blackberry%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QNX_external_memory_screen_buffer%20extension%2A)
- Aaron Ruby [\[GitHub\]aruby-blackberry](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QNX_external_memory_screen_buffer%5D%20%40aruby-blackberry%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QNX_external_memory_screen_buffer%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-05-17

**IP Status**

No known IP claims.

**Contributors**

- Mike Gorchak, QNX / Blackberry Limited
- Aaron Ruby, QNX / Blackberry Limited

## [](#_description)Description

This extension enables an application to import QNX Screen `_screen_buffer` objects created outside of the Vulkan device into Vulkan memory objects, where they can be bound to images and buffers.

Some `_screen_buffer` images have implementation-defined *external formats* that **may** not correspond to Vulkan formats. Sampler Y′CBCR conversion **can** be used to sample from these images and convert them to a known color space.

`_screen_buffer` is strongly typed, so naming the handle type is redundant. The internal layout and therefore size of a `_screen_buffer` image may depend on native usage flags that do not have corresponding Vulkan counterparts.

## [](#_new_commands)New Commands

- [vkGetScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetScreenBufferPropertiesQNX.html)

## [](#_new_structures)New Structures

- [VkScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenBufferPropertiesQNX.html)
- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html), [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html):
  
  - [VkExternalFormatQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatQNX.html)
- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html):
  
  - [VkImportScreenBufferInfoQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportScreenBufferInfoQNX.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX.html)
- Extending [VkScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenBufferPropertiesQNX.html):
  
  - [VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenBufferFormatPropertiesQNX.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_QNX_EXTERNAL_MEMORY_SCREEN_BUFFER_EXTENSION_NAME`
- `VK_QNX_EXTERNAL_MEMORY_SCREEN_BUFFER_SPEC_VERSION`
- Extending [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html):
  
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_SCREEN_BUFFER_BIT_QNX`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_EXTERNAL_FORMAT_QNX`
  - `VK_STRUCTURE_TYPE_IMPORT_SCREEN_BUFFER_INFO_QNX`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_MEMORY_SCREEN_BUFFER_FEATURES_QNX`
  - `VK_STRUCTURE_TYPE_SCREEN_BUFFER_FORMAT_PROPERTIES_QNX`
  - `VK_STRUCTURE_TYPE_SCREEN_BUFFER_PROPERTIES_QNX`

## [](#_issues)Issues

## [](#_version_history)Version History

- Revision 1, 2023-05-17 (Mike Gorchak)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QNX_external_memory_screen_buffer).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0