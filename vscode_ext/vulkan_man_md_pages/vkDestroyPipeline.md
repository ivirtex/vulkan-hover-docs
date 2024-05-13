# vkDestroyPipeline(3) Manual Page

## Name

vkDestroyPipeline - Destroy a pipeline object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a pipeline, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyPipeline(
    VkDevice                                    device,
    VkPipeline                                  pipeline,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that destroys the pipeline.

- `pipeline` is the handle of the pipeline to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkDestroyPipeline-pipeline-00765"
  id="VUID-vkDestroyPipeline-pipeline-00765"></a>
  VUID-vkDestroyPipeline-pipeline-00765  
  All submitted commands that refer to `pipeline` **must** have
  completed execution

- <a href="#VUID-vkDestroyPipeline-pipeline-00766"
  id="VUID-vkDestroyPipeline-pipeline-00766"></a>
  VUID-vkDestroyPipeline-pipeline-00766  
  If `VkAllocationCallbacks` were provided when `pipeline` was created,
  a compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyPipeline-pipeline-00767"
  id="VUID-vkDestroyPipeline-pipeline-00767"></a>
  VUID-vkDestroyPipeline-pipeline-00767  
  If no `VkAllocationCallbacks` were provided when `pipeline` was
  created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyPipeline-device-parameter"
  id="VUID-vkDestroyPipeline-device-parameter"></a>
  VUID-vkDestroyPipeline-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyPipeline-pipeline-parameter"
  id="VUID-vkDestroyPipeline-pipeline-parameter"></a>
  VUID-vkDestroyPipeline-pipeline-parameter  
  If `pipeline` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `pipeline`
  **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handle

- <a href="#VUID-vkDestroyPipeline-pAllocator-parameter"
  id="VUID-vkDestroyPipeline-pAllocator-parameter"></a>
  VUID-vkDestroyPipeline-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyPipeline-pipeline-parent"
  id="VUID-vkDestroyPipeline-pipeline-parent"></a>
  VUID-vkDestroyPipeline-pipeline-parent  
  If `pipeline` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `pipeline` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyPipeline"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
