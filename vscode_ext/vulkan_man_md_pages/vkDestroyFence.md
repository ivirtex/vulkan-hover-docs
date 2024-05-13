# vkDestroyFence(3) Manual Page

## Name

vkDestroyFence - Destroy a fence object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a fence, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyFence(
    VkDevice                                    device,
    VkFence                                     fence,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the fence.

- `fence` is the handle of the fence to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyFence-fence-01120"
  id="VUID-vkDestroyFence-fence-01120"></a>
  VUID-vkDestroyFence-fence-01120  
  All <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#devsandqueues-submission"
  target="_blank" rel="noopener">queue submission</a> commands that
  refer to `fence` **must** have completed execution

- <a href="#VUID-vkDestroyFence-fence-01121"
  id="VUID-vkDestroyFence-fence-01121"></a>
  VUID-vkDestroyFence-fence-01121  
  If `VkAllocationCallbacks` were provided when `fence` was created, a
  compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyFence-fence-01122"
  id="VUID-vkDestroyFence-fence-01122"></a>
  VUID-vkDestroyFence-fence-01122  
  If no `VkAllocationCallbacks` were provided when `fence` was created,
  `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyFence-device-parameter"
  id="VUID-vkDestroyFence-device-parameter"></a>
  VUID-vkDestroyFence-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyFence-fence-parameter"
  id="VUID-vkDestroyFence-fence-parameter"></a>
  VUID-vkDestroyFence-fence-parameter  
  If `fence` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `fence`
  **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html) handle

- <a href="#VUID-vkDestroyFence-pAllocator-parameter"
  id="VUID-vkDestroyFence-pAllocator-parameter"></a>
  VUID-vkDestroyFence-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyFence-fence-parent"
  id="VUID-vkDestroyFence-fence-parent"></a>
  VUID-vkDestroyFence-fence-parent  
  If `fence` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `fence` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyFence"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
