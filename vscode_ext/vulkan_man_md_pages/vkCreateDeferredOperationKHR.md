# vkCreateDeferredOperationKHR(3) Manual Page

## Name

vkCreateDeferredOperationKHR - Create a deferred operation handle



## <a href="#_c_specification" class="anchor"></a>C Specification

To construct the tracking object for a deferred command, call:

``` c
// Provided by VK_KHR_deferred_host_operations
VkResult vkCreateDeferredOperationKHR(
    VkDevice                                    device,
    const VkAllocationCallbacks*                pAllocator,
    VkDeferredOperationKHR*                     pDeferredOperation);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns `operation`.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pDeferredOperation` is a pointer to a handle in which the created
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkCreateDeferredOperationKHR-device-parameter"
  id="VUID-vkCreateDeferredOperationKHR-device-parameter"></a>
  VUID-vkCreateDeferredOperationKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateDeferredOperationKHR-pAllocator-parameter"
  id="VUID-vkCreateDeferredOperationKHR-pAllocator-parameter"></a>
  VUID-vkCreateDeferredOperationKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a
  href="#VUID-vkCreateDeferredOperationKHR-pDeferredOperation-parameter"
  id="VUID-vkCreateDeferredOperationKHR-pDeferredOperation-parameter"></a>
  VUID-vkCreateDeferredOperationKHR-pDeferredOperation-parameter  
  `pDeferredOperation` **must** be a valid pointer to a
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_deferred_host_operations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_deferred_host_operations.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateDeferredOperationKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
