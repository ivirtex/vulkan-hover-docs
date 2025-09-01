# vkGetDeviceMicromapCompatibilityEXT(3) Manual Page

## Name

vkGetDeviceMicromapCompatibilityEXT - Check if a serialized micromap is compatible with the current device



## [](#_c_specification)C Specification

To check if a serialized micromap is compatible with the current device call:

```c++
// Provided by VK_EXT_opacity_micromap
void vkGetDeviceMicromapCompatibilityEXT(
    VkDevice                                    device,
    const VkMicromapVersionInfoEXT*             pVersionInfo,
    VkAccelerationStructureCompatibilityKHR*    pCompatibility);
```

## [](#_parameters)Parameters

- `device` is the device to check the version against.
- `pVersionInfo` is a pointer to a [VkMicromapVersionInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapVersionInfoEXT.html) structure specifying version information to check against the device.
- `pCompatibility` is a pointer to a [VkAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCompatibilityKHR.html) value in which compatibility information is returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetDeviceMicromapCompatibilityEXT-micromap-07551)VUID-vkGetDeviceMicromapCompatibilityEXT-micromap-07551  
  The [`micromap`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-micromap) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetDeviceMicromapCompatibilityEXT-device-parameter)VUID-vkGetDeviceMicromapCompatibilityEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDeviceMicromapCompatibilityEXT-pVersionInfo-parameter)VUID-vkGetDeviceMicromapCompatibilityEXT-pVersionInfo-parameter  
  `pVersionInfo` **must** be a valid pointer to a valid [VkMicromapVersionInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapVersionInfoEXT.html) structure
- [](#VUID-vkGetDeviceMicromapCompatibilityEXT-pCompatibility-parameter)VUID-vkGetDeviceMicromapCompatibilityEXT-pCompatibility-parameter  
  `pCompatibility` **must** be a valid pointer to a [VkAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCompatibilityKHR.html) value

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCompatibilityKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkMicromapVersionInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapVersionInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDeviceMicromapCompatibilityEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0