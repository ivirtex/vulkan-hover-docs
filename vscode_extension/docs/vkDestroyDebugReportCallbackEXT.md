# vkDestroyDebugReportCallbackEXT(3) Manual Page

## Name

vkDestroyDebugReportCallbackEXT - Destroy a debug report callback object



## [](#_c_specification)C Specification

To destroy a `VkDebugReportCallbackEXT` object, call:

```c++
// Provided by VK_EXT_debug_report
void vkDestroyDebugReportCallbackEXT(
    VkInstance                                  instance,
    VkDebugReportCallbackEXT                    callback,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `instance` is the instance where the callback was created.
- `callback` is the [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackEXT.html) object to destroy. `callback` is an externally synchronized object and **must** not be used on more than one thread at a time. This means that `vkDestroyDebugReportCallbackEXT` **must** not be called when a callback is active.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyDebugReportCallbackEXT-instance-01242)VUID-vkDestroyDebugReportCallbackEXT-instance-01242  
  If `VkAllocationCallbacks` were provided when `callback` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyDebugReportCallbackEXT-instance-01243)VUID-vkDestroyDebugReportCallbackEXT-instance-01243  
  If no `VkAllocationCallbacks` were provided when `callback` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyDebugReportCallbackEXT-instance-parameter)VUID-vkDestroyDebugReportCallbackEXT-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkDestroyDebugReportCallbackEXT-callback-parameter)VUID-vkDestroyDebugReportCallbackEXT-callback-parameter  
  If `callback` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `callback` **must** be a valid [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackEXT.html) handle
- [](#VUID-vkDestroyDebugReportCallbackEXT-pAllocator-parameter)VUID-vkDestroyDebugReportCallbackEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyDebugReportCallbackEXT-callback-parent)VUID-vkDestroyDebugReportCallbackEXT-callback-parent  
  If `callback` is a valid handle, it **must** have been created, allocated, or retrieved from `instance`

Host Synchronization

- Host access to `callback` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_EXT\_debug\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_report.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackEXT.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyDebugReportCallbackEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0