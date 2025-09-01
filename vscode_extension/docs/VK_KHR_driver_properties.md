# VK\_KHR\_driver\_properties(3) Manual Page

## Name

VK\_KHR\_driver\_properties - device extension



## [](#_registered_extension_number)Registered Extension Number

197

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Daniel Rakos [\[GitHub\]drakos-amd](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_driver_properties%5D%20%40drakos-amd%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_driver_properties%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2018-04-11

**IP Status**

No known IP claims.

**Contributors**

- Baldur Karlsson
- Matthaeus G. Chajdas, AMD
- Piers Daniell, NVIDIA
- Alexander Galazin, Arm
- Jesse Hall, Google
- Daniel Rakos, AMD

## [](#_description)Description

This extension provides a new physical device query which allows retrieving information about the driver implementation, allowing applications to determine which physical device corresponds to which particular vendor’s driver, and which conformance test suite version the driver implementation is compliant with.

## [](#_promotion_to_vulkan_1_2)Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with the KHR suffix omitted. The original type, enum, and command names are still available as aliases of the core functionality.

## [](#_new_structures)New Structures

- [VkConformanceVersionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConformanceVersionKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceDriverPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDriverPropertiesKHR.html)

## [](#_new_enums)New Enums

- [VkDriverIdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDriverIdKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_DRIVER_PROPERTIES_EXTENSION_NAME`
- `VK_KHR_DRIVER_PROPERTIES_SPEC_VERSION`
- `VK_MAX_DRIVER_INFO_SIZE_KHR`
- `VK_MAX_DRIVER_NAME_SIZE_KHR`
- Extending [VkDriverId](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDriverId.html):
  
  - `VK_DRIVER_ID_AMD_OPEN_SOURCE_KHR`
  - `VK_DRIVER_ID_AMD_PROPRIETARY_KHR`
  - `VK_DRIVER_ID_ARM_PROPRIETARY_KHR`
  - `VK_DRIVER_ID_BROADCOM_PROPRIETARY_KHR`
  - `VK_DRIVER_ID_GGP_PROPRIETARY_KHR`
  - `VK_DRIVER_ID_GOOGLE_SWIFTSHADER_KHR`
  - `VK_DRIVER_ID_IMAGINATION_PROPRIETARY_KHR`
  - `VK_DRIVER_ID_INTEL_OPEN_SOURCE_MESA_KHR`
  - `VK_DRIVER_ID_INTEL_PROPRIETARY_WINDOWS_KHR`
  - `VK_DRIVER_ID_MESA_RADV_KHR`
  - `VK_DRIVER_ID_NVIDIA_PROPRIETARY_KHR`
  - `VK_DRIVER_ID_QUALCOMM_PROPRIETARY_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DRIVER_PROPERTIES_KHR`

## [](#_version_history)Version History

- Revision 1, 2018-04-11 (Daniel Rakos)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_driver_properties).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0