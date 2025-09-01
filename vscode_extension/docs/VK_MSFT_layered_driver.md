# VK\_MSFT\_layered\_driver(3) Manual Page

## Name

VK\_MSFT\_layered\_driver - device extension



## [](#_registered_extension_number)Registered Extension Number

531

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Jesse Natalie [\[GitHub\]jenatali](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_MSFT_layered_driver%5D%20%40jenatali%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_MSFT_layered_driver%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_MSFT\_layered\_driver](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_MSFT_layered_driver.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-06-21

**IP Status**

No known IP claims.

**Contributors**

- Jesse Natalie, Microsoft

## [](#_description)Description

This extension adds new physical device properties to allow applications and the Vulkan ICD loader to understand when a physical device is implemented as a layered driver on top of another underlying API.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceLayeredDriverPropertiesMSFT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredDriverPropertiesMSFT.html)

## [](#_new_enums)New Enums

- [VkLayeredDriverUnderlyingApiMSFT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkLayeredDriverUnderlyingApiMSFT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_MSFT_LAYERED_DRIVER_EXTENSION_NAME`
- `VK_MSFT_LAYERED_DRIVER_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_LAYERED_DRIVER_PROPERTIES_MSFT`

## [](#_examples)Examples

None.

## [](#_version_history)Version History

- Revision 1, 2023-06-21 (Jesse Natalie)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_MSFT_layered_driver).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0