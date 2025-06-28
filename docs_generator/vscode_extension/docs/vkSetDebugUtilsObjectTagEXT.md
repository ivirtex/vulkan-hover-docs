# vkSetDebugUtilsObjectTagEXT(3) Manual Page

## Name

vkSetDebugUtilsObjectTagEXT - Attach arbitrary data to an object



## [](#_c_specification)C Specification

```c++
// Provided by VK_EXT_debug_utils
VkResult vkSetDebugUtilsObjectTagEXT(
    VkDevice                                    device,
    const VkDebugUtilsObjectTagInfoEXT*         pTagInfo);
```

## [](#_parameters)Parameters

- `device` is the device that created the object.
- `pTagInfo` is a pointer to a [VkDebugUtilsObjectTagInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsObjectTagInfoEXT.html) structure specifying parameters of the tag to attach to the object.

## [](#_description)Description

Valid Usage

- [](#VUID-vkSetDebugUtilsObjectTagEXT-pNameInfo-07875)VUID-vkSetDebugUtilsObjectTagEXT-pNameInfo-07875  
  If `pNameInfo->objectHandle` is the valid handle of an instance-level object, the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) identified by `device` **must** be a descendent of the same [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) as the object identified by `pNameInfo->objectHandle`
- [](#VUID-vkSetDebugUtilsObjectTagEXT-pNameInfo-07876)VUID-vkSetDebugUtilsObjectTagEXT-pNameInfo-07876  
  If `pNameInfo->objectHandle` is the valid handle of a physical-device-level object, the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) identified by `device` **must** be a descendant of the same [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) as the object identified by `pNameInfo->objectHandle`
- [](#VUID-vkSetDebugUtilsObjectTagEXT-pNameInfo-07877)VUID-vkSetDebugUtilsObjectTagEXT-pNameInfo-07877  
  If `pNameInfo->objectHandle` is the valid handle of a device-level object, that object **must** be a descendent of the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) identified by `device`

Valid Usage (Implicit)

- [](#VUID-vkSetDebugUtilsObjectTagEXT-device-parameter)VUID-vkSetDebugUtilsObjectTagEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkSetDebugUtilsObjectTagEXT-pTagInfo-parameter)VUID-vkSetDebugUtilsObjectTagEXT-pTagInfo-parameter  
  `pTagInfo` **must** be a valid pointer to a valid [VkDebugUtilsObjectTagInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsObjectTagInfoEXT.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_EXT\_debug\_utils](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_utils.html), [VkDebugUtilsObjectTagInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsObjectTagInfoEXT.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkSetDebugUtilsObjectTagEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0