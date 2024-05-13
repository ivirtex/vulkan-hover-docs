# vkGetDeviceMicromapCompatibilityEXT(3) Manual Page

## Name

vkGetDeviceMicromapCompatibilityEXT - Check if a serialized micromap is
compatible with the current device



## <a href="#_c_specification" class="anchor"></a>C Specification

To check if a serialized micromap is compatible with the current device
call:

``` c
// Provided by VK_EXT_opacity_micromap
void vkGetDeviceMicromapCompatibilityEXT(
    VkDevice                                    device,
    const VkMicromapVersionInfoEXT*             pVersionInfo,
    VkAccelerationStructureCompatibilityKHR*    pCompatibility);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device to check the version against.

- `pVersionInfo` is a pointer to a
  [VkMicromapVersionInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapVersionInfoEXT.html) structure
  specifying version information to check against the device.

- `pCompatibility` is a pointer to a
  [VkAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCompatibilityKHR.html)
  value in which compatibility information is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkGetDeviceMicromapCompatibilityEXT-micromap-07551"
  id="VUID-vkGetDeviceMicromapCompatibilityEXT-micromap-07551"></a>
  VUID-vkGetDeviceMicromapCompatibilityEXT-micromap-07551  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-micromap"
  target="_blank" rel="noopener"><code>micromap</code></a> feature
  **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkGetDeviceMicromapCompatibilityEXT-device-parameter"
  id="VUID-vkGetDeviceMicromapCompatibilityEXT-device-parameter"></a>
  VUID-vkGetDeviceMicromapCompatibilityEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetDeviceMicromapCompatibilityEXT-pVersionInfo-parameter"
  id="VUID-vkGetDeviceMicromapCompatibilityEXT-pVersionInfo-parameter"></a>
  VUID-vkGetDeviceMicromapCompatibilityEXT-pVersionInfo-parameter  
  `pVersionInfo` **must** be a valid pointer to a valid
  [VkMicromapVersionInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapVersionInfoEXT.html) structure

- <a
  href="#VUID-vkGetDeviceMicromapCompatibilityEXT-pCompatibility-parameter"
  id="VUID-vkGetDeviceMicromapCompatibilityEXT-pCompatibility-parameter"></a>
  VUID-vkGetDeviceMicromapCompatibilityEXT-pCompatibility-parameter  
  `pCompatibility` **must** be a valid pointer to a
  [VkAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCompatibilityKHR.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCompatibilityKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkMicromapVersionInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapVersionInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeviceMicromapCompatibilityEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
