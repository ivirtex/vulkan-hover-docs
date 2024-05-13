# vkGetDeviceAccelerationStructureCompatibilityKHR(3) Manual Page

## Name

vkGetDeviceAccelerationStructureCompatibilityKHR - Check if a serialized
acceleration structure is compatible with the current device



## <a href="#_c_specification" class="anchor"></a>C Specification

To check if a serialized acceleration structure is compatible with the
current device call:

``` c
// Provided by VK_KHR_acceleration_structure
void vkGetDeviceAccelerationStructureCompatibilityKHR(
    VkDevice                                    device,
    const VkAccelerationStructureVersionInfoKHR* pVersionInfo,
    VkAccelerationStructureCompatibilityKHR*    pCompatibility);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device to check the version against.

- `pVersionInfo` is a pointer to a
  [VkAccelerationStructureVersionInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureVersionInfoKHR.html)
  structure specifying version information to check against the device.

- `pCompatibility` is a pointer to a
  [VkAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCompatibilityKHR.html)
  value in which compatibility information is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-accelerationStructure-08928"
  id="VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-accelerationStructure-08928"></a>
  VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-accelerationStructure-08928  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-accelerationStructure"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceAccelerationStructureFeaturesKHR</code>::<code>accelerationStructure</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-device-parameter"
  id="VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-device-parameter"></a>
  VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-pVersionInfo-parameter"
  id="VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-pVersionInfo-parameter"></a>
  VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-pVersionInfo-parameter  
  `pVersionInfo` **must** be a valid pointer to a valid
  [VkAccelerationStructureVersionInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureVersionInfoKHR.html)
  structure

- <a
  href="#VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-pCompatibility-parameter"
  id="VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-pCompatibility-parameter"></a>
  VUID-vkGetDeviceAccelerationStructureCompatibilityKHR-pCompatibility-parameter  
  `pCompatibility` **must** be a valid pointer to a
  [VkAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCompatibilityKHR.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCompatibilityKHR.html),
[VkAccelerationStructureVersionInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureVersionInfoKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeviceAccelerationStructureCompatibilityKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
