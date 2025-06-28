# vkSetDebugUtilsObjectNameEXT(3) Manual Page

## Name

vkSetDebugUtilsObjectNameEXT - Give an application-defined name to an object



## [](#_c_specification)C Specification

An object can be given an application-defined name by calling:

```c++
// Provided by VK_EXT_debug_utils
VkResult vkSetDebugUtilsObjectNameEXT(
    VkDevice                                    device,
    const VkDebugUtilsObjectNameInfoEXT*        pNameInfo);
```

## [](#_parameters)Parameters

- `device` is the device that is associated with the named object passed in via `objectHandle`.
- `pNameInfo` is a pointer to a [VkDebugUtilsObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsObjectNameInfoEXT.html) structure specifying parameters of the name to set on the object.

## [](#_description)Description

Valid Usage

- [](#VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-02587)VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-02587  
  `pNameInfo->objectType` **must** not be `VK_OBJECT_TYPE_UNKNOWN`
- [](#VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-02588)VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-02588  
  `pNameInfo->objectHandle` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-07872)VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-07872  
  If `pNameInfo->objectHandle` is the valid handle of an instance-level object, the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) identified by `device` **must** be a descendent of the same [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) as the object identified by `pNameInfo->objectHandle`
- [](#VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-07873)VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-07873  
  If `pNameInfo->objectHandle` is the valid handle of a physical-device-level object, the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) identified by `device` **must** be a descendant of the same [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) as the object identified by `pNameInfo->objectHandle`
- [](#VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-07874)VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-07874  
  If `pNameInfo->objectHandle` is the valid handle of a device-level object, that object **must** be a descendent of the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) identified by `device`

Valid Usage (Implicit)

- [](#VUID-vkSetDebugUtilsObjectNameEXT-device-parameter)VUID-vkSetDebugUtilsObjectNameEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-parameter)VUID-vkSetDebugUtilsObjectNameEXT-pNameInfo-parameter  
  `pNameInfo` **must** be a valid pointer to a valid [VkDebugUtilsObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsObjectNameInfoEXT.html) structure

Host Synchronization

- Host access to `pNameInfo->objectHandle` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_EXT\_debug\_utils](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_utils.html), [VkDebugUtilsObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsObjectNameInfoEXT.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkSetDebugUtilsObjectNameEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0