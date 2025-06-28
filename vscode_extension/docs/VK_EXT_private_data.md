# VK\_EXT\_private\_data(3) Manual Page

## Name

VK\_EXT\_private\_data - device extension



## [](#_registered_extension_number)Registered Extension Number

296

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_contact)Contact

- Matthew Rusch [\[GitHub\]mattruschnv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_private_data%5D%20%40mattruschnv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_private_data%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-03-25

**IP Status**

No known IP claims.

**Contributors**

- Matthew Rusch, NVIDIA
- Nuno Subtil, NVIDIA
- Piers Daniell, NVIDIA
- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension is a device extension which enables attaching arbitrary payloads to Vulkan objects. It introduces the idea of private data slots as a means of storing a 64-bit unsigned integer of application-defined data. Private data slots can be created or destroyed any time an associated device is available. Private data slots can be reserved at device creation time, and limiting use to the amount reserved will allow the extension to exhibit better performance characteristics.

## [](#_new_object_types)New Object Types

- [VkPrivateDataSlotEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlotEXT.html)

## [](#_new_commands)New Commands

- [vkCreatePrivateDataSlotEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreatePrivateDataSlotEXT.html)
- [vkDestroyPrivateDataSlotEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyPrivateDataSlotEXT.html)
- [vkGetPrivateDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPrivateDataEXT.html)
- [vkSetPrivateDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetPrivateDataEXT.html)

## [](#_new_structures)New Structures

- [VkPrivateDataSlotCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlotCreateInfoEXT.html)
- Extending [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkDevicePrivateDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevicePrivateDataCreateInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePrivateDataFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePrivateDataFeaturesEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkPrivateDataSlotCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlotCreateFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_PRIVATE_DATA_EXTENSION_NAME`
- `VK_EXT_PRIVATE_DATA_SPEC_VERSION`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_PRIVATE_DATA_SLOT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DEVICE_PRIVATE_DATA_CREATE_INFO_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRIVATE_DATA_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PRIVATE_DATA_SLOT_CREATE_INFO_EXT`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

Vulkan APIs in this extension are included in core Vulkan 1.3, with the EXT suffix omitted. External interactions defined by this extension, such as SPIR-V token names, retain their original names. The original Vulkan API names are still available as aliases of the core functionality.

## [](#_examples)Examples

- In progress

## [](#_issues)Issues

(1) If I have to create a [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html) to store and retrieve data on an object, how does this extension help me? Will I not need to store the [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html) mapping with each object, and if I am doing that, I might as well just store the original data!

**RESOLVED:** The [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html) can be thought of as an opaque index into storage that is reserved in each object. That is, you can use the same [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html) with each object for a specific piece of information. For example, if a layer wishes to track per-object information, the layer only needs to allocate one [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html) per device and it can use that private data slot for all of the device’s child objects. This allows multiple layers to store private data without conflicting with each other’s and/or the application’s private data.

(2) What if I need to store more than 64-bits of information per object?

**RESOLVED:** The data that you store per object could be a pointer to another object or structure of your own allocation.

## [](#_version_history)Version History

- Revision 1, 2020-01-15 (Matthew Rusch)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_private_data)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0