# VK\_KHR\_present\_id(3) Manual Page

## Name

VK\_KHR\_present\_id - device extension



## [](#_registered_extension_number)Registered Extension Number

295

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html)  
and  
[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Keith Packard [\[GitHub\]keithp](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_present_id%5D%20%40keithp%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_present_id%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-05-15

**IP Status**

No known IP claims.

**Contributors**

- Keith Packard, Valve
- Ian Elliott, Google
- Alon Or-bach, Samsung

## [](#_description)Description

This device extension allows an application that uses the `VK_KHR_swapchain` extension to provide an identifier for present operations on a swapchain. An application **can** use this to reference specific present operations in other extensions.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePresentIdFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePresentIdFeaturesKHR.html)
- Extending [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html):
  
  - [VkPresentIdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentIdKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_PRESENT_ID_EXTENSION_NAME`
- `VK_KHR_PRESENT_ID_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PRESENT_ID_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PRESENT_ID_KHR`

## [](#_issues)Issues

None.

## [](#_examples)Examples

## [](#_version_history)Version History

- Revision 1, 2019-05-15 (Keith Packard)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_present_id)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0