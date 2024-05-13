# vkDestroyDebugUtilsMessengerEXT(3) Manual Page

## Name

vkDestroyDebugUtilsMessengerEXT - Destroy a debug messenger object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a `VkDebugUtilsMessengerEXT` object, call:

``` c
// Provided by VK_EXT_debug_utils
void vkDestroyDebugUtilsMessengerEXT(
    VkInstance                                  instance,
    VkDebugUtilsMessengerEXT                    messenger,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `instance` is the instance where the callback was created.

- `messenger` is the
  [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerEXT.html) object to
  destroy. `messenger` is an externally synchronized object and **must**
  not be used on more than one thread at a time. This means that
  `vkDestroyDebugUtilsMessengerEXT` **must** not be called when a
  callback is active.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyDebugUtilsMessengerEXT-messenger-01915"
  id="VUID-vkDestroyDebugUtilsMessengerEXT-messenger-01915"></a>
  VUID-vkDestroyDebugUtilsMessengerEXT-messenger-01915  
  If `VkAllocationCallbacks` were provided when `messenger` was created,
  a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyDebugUtilsMessengerEXT-messenger-01916"
  id="VUID-vkDestroyDebugUtilsMessengerEXT-messenger-01916"></a>
  VUID-vkDestroyDebugUtilsMessengerEXT-messenger-01916  
  If no `VkAllocationCallbacks` were provided when `messenger` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyDebugUtilsMessengerEXT-instance-parameter"
  id="VUID-vkDestroyDebugUtilsMessengerEXT-instance-parameter"></a>
  VUID-vkDestroyDebugUtilsMessengerEXT-instance-parameter  
  `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html) handle

- <a href="#VUID-vkDestroyDebugUtilsMessengerEXT-messenger-parameter"
  id="VUID-vkDestroyDebugUtilsMessengerEXT-messenger-parameter"></a>
  VUID-vkDestroyDebugUtilsMessengerEXT-messenger-parameter  
  If `messenger` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `messenger` **must** be a valid
  [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerEXT.html) handle

- <a href="#VUID-vkDestroyDebugUtilsMessengerEXT-pAllocator-parameter"
  id="VUID-vkDestroyDebugUtilsMessengerEXT-pAllocator-parameter"></a>
  VUID-vkDestroyDebugUtilsMessengerEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyDebugUtilsMessengerEXT-messenger-parent"
  id="VUID-vkDestroyDebugUtilsMessengerEXT-messenger-parent"></a>
  VUID-vkDestroyDebugUtilsMessengerEXT-messenger-parent  
  If `messenger` is a valid handle, it **must** have been created,
  allocated, or retrieved from `instance`

Host Synchronization

- Host access to `messenger` **must** be externally synchronized

The application **must** ensure that
[vkDestroyDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDebugUtilsMessengerEXT.html)
is not executed in parallel with any Vulkan command that is also called
with `instance` or child of `instance` as the dispatchable argument.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_utils](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_utils.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerEXT.html),
[VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyDebugUtilsMessengerEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
