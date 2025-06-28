# vkGetDeviceAccelerationStructureCompatibilityKHR(3) Manual Page

## Name

vkGetDeviceAccelerationStructureCompatibilityKHR - Check if a serialized acceleration structure is compatible with the current device



## [](#_c_specification)C Specification

To check if a serialized acceleration structure is compatible with the current device call:

```c++
// Provided by VK_KHR_acceleration_structure
void vkGetDeviceAccelerationStructureCompatibilityKHR(
    VkDevice                                    device,
    const VkAccelerationStructureVersionInfoKHR* pVersionInfo,
    VkAccelerationStructureCompatibilityKHR*    pCompatibility);
```

## [](#_parameters)Parameters

- `device` is the device to check the version against.
- `pVersionInfo` is a pointer to a [VkAccelerationStructureVersionInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureVersionInfoKHR.html) structure specifying version information to check against the device.
- `pCompatibility` is a pointer to a [VkAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCompatibilityKHR.html) value in which compatibility information is returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-accelerationStructure-08928)VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-accelerationStructure-08928  
  The [`VkPhysicalDeviceAccelerationStructureFeaturesKHR`::`accelerationStructure`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-accelerationStructure) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-device-parameter)VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-pVersionInfo-parameter)VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-pVersionInfo-parameter  
  `pVersionInfo` **must** be a valid pointer to a valid [VkAccelerationStructureVersionInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureVersionInfoKHR.html) structure
- [](#VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-pCompatibility-parameter)VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-pCompatibility-parameter  
  `pCompatibility` **must** be a valid pointer to a [VkAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCompatibilityKHR.html) value

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCompatibilityKHR.html), [VkAccelerationStructureVersionInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureVersionInfoKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDeviceAccelerationStructureCompatibilityKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0