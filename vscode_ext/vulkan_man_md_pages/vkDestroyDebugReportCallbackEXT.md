# vkDestroyDebugReportCallbackEXT(3) Manual Page

## Name

vkDestroyDebugReportCallbackEXT - Destroy a debug report callback object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a `VkDebugReportCallbackEXT` object, call:

``` c
// Provided by VK_EXT_debug_report
void vkDestroyDebugReportCallbackEXT(
    VkInstance                                  instance,
    VkDebugReportCallbackEXT                    callback,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is the instance where the callback was created.

- `callback` is the
  [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackEXT.html) object to
  destroy. `callback` is an externally synchronized object and **must**
  not be used on more than one thread at a time. This means that
  `vkDestroyDebugReportCallbackEXT` **must** not be called when a
  callback is active.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyDebugReportCallbackEXT-instance-01242"
  id="VUID-vkDestroyDebugReportCallbackEXT-instance-01242"></a>
  VUID-vkDestroyDebugReportCallbackEXT-instance-01242  
  If `VkAllocationCallbacks` were provided when `callback` was created,
  a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyDebugReportCallbackEXT-instance-01243"
  id="VUID-vkDestroyDebugReportCallbackEXT-instance-01243"></a>
  VUID-vkDestroyDebugReportCallbackEXT-instance-01243  
  If no `VkAllocationCallbacks` were provided when `callback` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyDebugReportCallbackEXT-instance-parameter"
  id="VUID-vkDestroyDebugReportCallbackEXT-instance-parameter"></a>
  VUID-vkDestroyDebugReportCallbackEXT-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a href="#VUID-vkDestroyDebugReportCallbackEXT-callback-parameter"
  id="VUID-vkDestroyDebugReportCallbackEXT-callback-parameter"></a>
  VUID-vkDestroyDebugReportCallbackEXT-callback-parameter  
  If `callback` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `callback`
  **must** be a valid
  [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackEXT.html) handle

- <a href="#VUID-vkDestroyDebugReportCallbackEXT-pAllocator-parameter"
  id="VUID-vkDestroyDebugReportCallbackEXT-pAllocator-parameter"></a>
  VUID-vkDestroyDebugReportCallbackEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyDebugReportCallbackEXT-callback-parent"
  id="VUID-vkDestroyDebugReportCallbackEXT-callback-parent"></a>
  VUID-vkDestroyDebugReportCallbackEXT-callback-parent  
  If `callback` is a valid handle, it **must** have been created,
  allocated, or retrieved from `instance`

Host Synchronization

- Host access to `callback` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_report.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportCallbackEXT.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyDebugReportCallbackEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
