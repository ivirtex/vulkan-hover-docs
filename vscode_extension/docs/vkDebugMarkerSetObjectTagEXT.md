# vkDebugMarkerSetObjectTagEXT(3) Manual Page

## Name

vkDebugMarkerSetObjectTagEXT - Attach arbitrary data to an object



## [](#_c_specification)C Specification

In addition to setting a name for an object, debugging and validation layers may have uses for additional binary data on a per-object basis that has no other place in the Vulkan API. For example, a `VkShaderModule` could have additional debugging data attached to it to aid in offline shader tracing. To attach data to an object, call:

```c++
// Provided by VK_EXT_debug_marker
VkResult vkDebugMarkerSetObjectTagEXT(
    VkDevice                                    device,
    const VkDebugMarkerObjectTagInfoEXT*        pTagInfo);
```

## [](#_parameters)Parameters

- `device` is the device that created the object.
- `pTagInfo` is a pointer to a [VkDebugMarkerObjectTagInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugMarkerObjectTagInfoEXT.html) structure specifying the parameters of the tag to attach to the object.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkDebugMarkerSetObjectTagEXT-device-parameter)VUID-vkDebugMarkerSetObjectTagEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDebugMarkerSetObjectTagEXT-pTagInfo-parameter)VUID-vkDebugMarkerSetObjectTagEXT-pTagInfo-parameter  
  `pTagInfo` **must** be a valid pointer to a valid [VkDebugMarkerObjectTagInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugMarkerObjectTagInfoEXT.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_debug\_marker](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_marker.html), [VkDebugMarkerObjectTagInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugMarkerObjectTagInfoEXT.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDebugMarkerSetObjectTagEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0