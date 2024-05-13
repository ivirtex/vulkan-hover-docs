# vkDestroySampler(3) Manual Page

## Name

vkDestroySampler - Destroy a sampler object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a sampler, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroySampler(
    VkDevice                                    device,
    VkSampler                                   sampler,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the sampler.

- `sampler` is the sampler to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroySampler-sampler-01082"
  id="VUID-vkDestroySampler-sampler-01082"></a>
  VUID-vkDestroySampler-sampler-01082  
  All submitted commands that refer to `sampler` **must** have completed
  execution

- <a href="#VUID-vkDestroySampler-sampler-01083"
  id="VUID-vkDestroySampler-sampler-01083"></a>
  VUID-vkDestroySampler-sampler-01083  
  If `VkAllocationCallbacks` were provided when `sampler` was created, a
  compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroySampler-sampler-01084"
  id="VUID-vkDestroySampler-sampler-01084"></a>
  VUID-vkDestroySampler-sampler-01084  
  If no `VkAllocationCallbacks` were provided when `sampler` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroySampler-device-parameter"
  id="VUID-vkDestroySampler-device-parameter"></a>
  VUID-vkDestroySampler-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroySampler-sampler-parameter"
  id="VUID-vkDestroySampler-sampler-parameter"></a>
  VUID-vkDestroySampler-sampler-parameter  
  If `sampler` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `sampler`
  **must** be a valid [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html) handle

- <a href="#VUID-vkDestroySampler-pAllocator-parameter"
  id="VUID-vkDestroySampler-pAllocator-parameter"></a>
  VUID-vkDestroySampler-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroySampler-sampler-parent"
  id="VUID-vkDestroySampler-sampler-parent"></a>
  VUID-vkDestroySampler-sampler-parent  
  If `sampler` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `sampler` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampler.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroySampler"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
