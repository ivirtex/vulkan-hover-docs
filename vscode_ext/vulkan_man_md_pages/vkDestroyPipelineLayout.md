# vkDestroyPipelineLayout(3) Manual Page

## Name

vkDestroyPipelineLayout - Destroy a pipeline layout object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a pipeline layout, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyPipelineLayout(
    VkDevice                                    device,
    VkPipelineLayout                            pipelineLayout,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the pipeline layout.

- `pipelineLayout` is the pipeline layout to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyPipelineLayout-pipelineLayout-00299"
  id="VUID-vkDestroyPipelineLayout-pipelineLayout-00299"></a>
  VUID-vkDestroyPipelineLayout-pipelineLayout-00299  
  If `VkAllocationCallbacks` were provided when `pipelineLayout` was
  created, a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyPipelineLayout-pipelineLayout-00300"
  id="VUID-vkDestroyPipelineLayout-pipelineLayout-00300"></a>
  VUID-vkDestroyPipelineLayout-pipelineLayout-00300  
  If no `VkAllocationCallbacks` were provided when `pipelineLayout` was
  created, `pAllocator` **must** be `NULL`

- <a href="#VUID-vkDestroyPipelineLayout-pipelineLayout-02004"
  id="VUID-vkDestroyPipelineLayout-pipelineLayout-02004"></a>
  VUID-vkDestroyPipelineLayout-pipelineLayout-02004  
  `pipelineLayout` **must** not have been passed to any `vkCmd*` command
  for any command buffers that are still in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
  target="_blank" rel="noopener">recording state</a> when
  `vkDestroyPipelineLayout` is called

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyPipelineLayout-device-parameter"
  id="VUID-vkDestroyPipelineLayout-device-parameter"></a>
  VUID-vkDestroyPipelineLayout-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyPipelineLayout-pipelineLayout-parameter"
  id="VUID-vkDestroyPipelineLayout-pipelineLayout-parameter"></a>
  VUID-vkDestroyPipelineLayout-pipelineLayout-parameter  
  If `pipelineLayout` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pipelineLayout` **must** be a valid
  [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) handle

- <a href="#VUID-vkDestroyPipelineLayout-pAllocator-parameter"
  id="VUID-vkDestroyPipelineLayout-pAllocator-parameter"></a>
  VUID-vkDestroyPipelineLayout-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyPipelineLayout-pipelineLayout-parent"
  id="VUID-vkDestroyPipelineLayout-pipelineLayout-parent"></a>
  VUID-vkDestroyPipelineLayout-pipelineLayout-parent  
  If `pipelineLayout` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `pipelineLayout` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyPipelineLayout"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
