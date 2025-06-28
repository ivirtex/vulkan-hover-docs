# vkCreateDebugUtilsMessengerEXT(3) Manual Page

## Name

vkCreateDebugUtilsMessengerEXT - Create a debug messenger object



## [](#_c_specification)C Specification

A debug messenger triggers a debug callback with a debug message when an event of interest occurs. To create a debug messenger which will trigger a debug callback, call:

```c++
// Provided by VK_EXT_debug_utils
VkResult vkCreateDebugUtilsMessengerEXT(
    VkInstance                                  instance,
    const VkDebugUtilsMessengerCreateInfoEXT*   pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDebugUtilsMessengerEXT*                   pMessenger);
```

## [](#_parameters)Parameters

- `instance` is the instance the messenger will be used with.
- `pCreateInfo` is a pointer to a [VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCreateInfoEXT.html) structure containing the callback pointer, as well as defining conditions under which this messenger will trigger the callback.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pMessenger` is a pointer to a [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerEXT.html) handle in which the created object is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateDebugUtilsMessengerEXT-instance-parameter)VUID-vkCreateDebugUtilsMessengerEXT-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkCreateDebugUtilsMessengerEXT-pCreateInfo-parameter)VUID-vkCreateDebugUtilsMessengerEXT-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCreateInfoEXT.html) structure
- [](#VUID-vkCreateDebugUtilsMessengerEXT-pAllocator-parameter)VUID-vkCreateDebugUtilsMessengerEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateDebugUtilsMessengerEXT-pMessenger-parameter)VUID-vkCreateDebugUtilsMessengerEXT-pMessenger-parameter  
  `pMessenger` **must** be a valid pointer to a [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerEXT.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`

The application **must** ensure that [vkCreateDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDebugUtilsMessengerEXT.html) is not executed in parallel with any Vulkan command that is also called with `instance` or child of `instance` as the dispatchable argument.

## [](#_see_also)See Also

[VK\_EXT\_debug\_utils](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_utils.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerCreateInfoEXT.html), [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerEXT.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateDebugUtilsMessengerEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0