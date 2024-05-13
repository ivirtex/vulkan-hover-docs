# vkCreateDebugUtilsMessengerEXT(3) Manual Page

## Name

vkCreateDebugUtilsMessengerEXT - Create a debug messenger object



## <a href="#_c_specification" class="anchor"></a>C Specification

A debug messenger triggers a debug callback with a debug message when an
event of interest occurs. To create a debug messenger which will trigger
a debug callback, call:

``` c
// Provided by VK_EXT_debug_utils
VkResult vkCreateDebugUtilsMessengerEXT(
    VkInstance                                  instance,
    const VkDebugUtilsMessengerCreateInfoEXT*   pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDebugUtilsMessengerEXT*                   pMessenger);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is the instance the messenger will be used with.

- `pCreateInfo` is a pointer to a
  [VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCreateInfoEXT.html)
  structure containing the callback pointer, as well as defining
  conditions under which this messenger will trigger the callback.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pMessenger` is a pointer to a
  [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerEXT.html) handle in
  which the created object is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkCreateDebugUtilsMessengerEXT-instance-parameter"
  id="VUID-vkCreateDebugUtilsMessengerEXT-instance-parameter"></a>
  VUID-vkCreateDebugUtilsMessengerEXT-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a href="#VUID-vkCreateDebugUtilsMessengerEXT-pCreateInfo-parameter"
  id="VUID-vkCreateDebugUtilsMessengerEXT-pCreateInfo-parameter"></a>
  VUID-vkCreateDebugUtilsMessengerEXT-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCreateInfoEXT.html)
  structure

- <a href="#VUID-vkCreateDebugUtilsMessengerEXT-pAllocator-parameter"
  id="VUID-vkCreateDebugUtilsMessengerEXT-pAllocator-parameter"></a>
  VUID-vkCreateDebugUtilsMessengerEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateDebugUtilsMessengerEXT-pMessenger-parameter"
  id="VUID-vkCreateDebugUtilsMessengerEXT-pMessenger-parameter"></a>
  VUID-vkCreateDebugUtilsMessengerEXT-pMessenger-parameter  
  `pMessenger` **must** be a valid pointer to a
  [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerEXT.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

The application **must** ensure that
[vkCreateDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDebugUtilsMessengerEXT.html) is
not executed in parallel with any Vulkan command that is also called
with `instance` or child of `instance` as the dispatchable argument.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDebugUtilsMessengerCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCreateInfoEXT.html),
[VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerEXT.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateDebugUtilsMessengerEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
