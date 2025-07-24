# vkDebugMarkerSetObjectNameEXT(3) Manual Page

## Name

vkDebugMarkerSetObjectNameEXT - Give an application-defined name to an object



## [](#_c_specification)C Specification

An object can be given an application-defined name by calling:

```c++
// Provided by VK_EXT_debug_marker
VkResult vkDebugMarkerSetObjectNameEXT(
    VkDevice                                    device,
    const VkDebugMarkerObjectNameInfoEXT*       pNameInfo);
```

## [](#_parameters)Parameters

- `device` is the device that created the object.
- `pNameInfo` is a pointer to a [VkDebugMarkerObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugMarkerObjectNameInfoEXT.html) structure specifying the parameters of the name to set on the object.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkDebugMarkerSetObjectNameEXT-device-parameter)VUID-vkDebugMarkerSetObjectNameEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDebugMarkerSetObjectNameEXT-pNameInfo-parameter)VUID-vkDebugMarkerSetObjectNameEXT-pNameInfo-parameter  
  `pNameInfo` **must** be a valid pointer to a valid [VkDebugMarkerObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugMarkerObjectNameInfoEXT.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_debug\_marker](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_marker.html), [VkDebugMarkerObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugMarkerObjectNameInfoEXT.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDebugMarkerSetObjectNameEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0