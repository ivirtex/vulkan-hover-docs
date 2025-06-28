# VK\_KHR\_zero\_initialize\_workgroup\_memory(3) Manual Page

## Name

VK\_KHR\_zero\_initialize\_workgroup\_memory - device extension



## [](#_registered_extension_number)Registered Extension Number

326

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_contact)Contact

- Alan Baker [\[GitHub\]alan-baker](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_zero_initialize_workgroup_memory%5D%20%40alan-baker%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_zero_initialize_workgroup_memory%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-11-18

**IP Status**

No known IP claims.

**Contributors**

- Alan Baker, Google
- Jeff Bolz, Nvidia
- Faith Ekstrand, Intel

## [](#_description)Description

This extension allows the use of a null constant initializer on shader Workgroup memory variables, allowing implementations to expose any special hardware or instructions they may have. Zero initialization is commonly used by applications running untrusted content (e.g. web browsers) as way of defeating memory-scraping attacks.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceZeroInitializeWorkgroupMemoryFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_ZERO_INITIALIZE_WORKGROUP_MEMORY_EXTENSION_NAME`
- `VK_KHR_ZERO_INITIALIZE_WORKGROUP_MEMORY_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_ZERO_INITIALIZE_WORKGROUP_MEMORY_FEATURES_KHR`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

Vulkan APIs in this extension are included in core Vulkan 1.3, with the KHR suffix omitted. External interactions defined by this extension, such as SPIR-V token names, retain their original names. The original Vulkan API names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 1, 2020-11-18 (Alan Baker)
  
  - Internal draft version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_zero_initialize_workgroup_memory)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0