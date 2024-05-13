# vkDestroyDeferredOperationKHR(3) Manual Page

## Name

vkDestroyDeferredOperationKHR - Destroy a deferred operation handle



## <a href="#_c_specification" class="anchor"></a>C Specification

When a deferred operation is completed, the application **can** destroy
the tracking object by calling:

``` c
// Provided by VK_KHR_deferred_host_operations
void vkDestroyDeferredOperationKHR(
    VkDevice                                    device,
    VkDeferredOperationKHR                      operation,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns `operation`.

- `operation` is the completed operation to be destroyed.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyDeferredOperationKHR-operation-03434"
  id="VUID-vkDestroyDeferredOperationKHR-operation-03434"></a>
  VUID-vkDestroyDeferredOperationKHR-operation-03434  
  If `VkAllocationCallbacks` were provided when `operation` was created,
  a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyDeferredOperationKHR-operation-03435"
  id="VUID-vkDestroyDeferredOperationKHR-operation-03435"></a>
  VUID-vkDestroyDeferredOperationKHR-operation-03435  
  If no `VkAllocationCallbacks` were provided when `operation` was
  created, `pAllocator` **must** be `NULL`

- <a href="#VUID-vkDestroyDeferredOperationKHR-operation-03436"
  id="VUID-vkDestroyDeferredOperationKHR-operation-03436"></a>
  VUID-vkDestroyDeferredOperationKHR-operation-03436  
  `operation` **must** be completed

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyDeferredOperationKHR-device-parameter"
  id="VUID-vkDestroyDeferredOperationKHR-device-parameter"></a>
  VUID-vkDestroyDeferredOperationKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyDeferredOperationKHR-operation-parameter"
  id="VUID-vkDestroyDeferredOperationKHR-operation-parameter"></a>
  VUID-vkDestroyDeferredOperationKHR-operation-parameter  
  If `operation` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `operation` **must** be a valid
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) handle

- <a href="#VUID-vkDestroyDeferredOperationKHR-pAllocator-parameter"
  id="VUID-vkDestroyDeferredOperationKHR-pAllocator-parameter"></a>
  VUID-vkDestroyDeferredOperationKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyDeferredOperationKHR-operation-parent"
  id="VUID-vkDestroyDeferredOperationKHR-operation-parent"></a>
  VUID-vkDestroyDeferredOperationKHR-operation-parent  
  If `operation` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `operation` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_deferred_host_operations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_deferred_host_operations.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyDeferredOperationKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
