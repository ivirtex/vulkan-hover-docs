# vkDestroyDebugUtilsMessengerEXT(3) Manual Page

## Name

vkDestroyDebugUtilsMessengerEXT - Destroy a debug messenger object



## [](#_c_specification)C Specification

To destroy a `VkDebugUtilsMessengerEXT` object, call:

```c++
// Provided by VK_EXT_debug_utils
void vkDestroyDebugUtilsMessengerEXT(
    VkInstance                                  instance,
    VkDebugUtilsMessengerEXT                    messenger,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `instance` is the instance where the callback was created.
- `messenger` is the [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerEXT.html) object to destroy. `messenger` is an externally synchronized object and **must** not be used on more than one thread at a time. This means that `vkDestroyDebugUtilsMessengerEXT` **must** not be called when a callback is active.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyDebugUtilsMessengerEXT-messenger-01915)VUID-vkDestroyDebugUtilsMessengerEXT-messenger-01915  
  If `VkAllocationCallbacks` were provided when `messenger` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyDebugUtilsMessengerEXT-messenger-01916)VUID-vkDestroyDebugUtilsMessengerEXT-messenger-01916  
  If no `VkAllocationCallbacks` were provided when `messenger` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyDebugUtilsMessengerEXT-instance-parameter)VUID-vkDestroyDebugUtilsMessengerEXT-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkDestroyDebugUtilsMessengerEXT-messenger-parameter)VUID-vkDestroyDebugUtilsMessengerEXT-messenger-parameter  
  If `messenger` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `messenger` **must** be a valid [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerEXT.html) handle
- [](#VUID-vkDestroyDebugUtilsMessengerEXT-pAllocator-parameter)VUID-vkDestroyDebugUtilsMessengerEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyDebugUtilsMessengerEXT-messenger-parent)VUID-vkDestroyDebugUtilsMessengerEXT-messenger-parent  
  If `messenger` is a valid handle, it **must** have been created, allocated, or retrieved from `instance`

Host Synchronization

- Host access to `messenger` **must** be externally synchronized

The application **must** ensure that [vkDestroyDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyDebugUtilsMessengerEXT.html) is not executed in parallel with any Vulkan command that is also called with `instance` or child of `instance` as the dispatchable argument.

## [](#_see_also)See Also

[VK\_EXT\_debug\_utils](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_utils.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerEXT.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyDebugUtilsMessengerEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0