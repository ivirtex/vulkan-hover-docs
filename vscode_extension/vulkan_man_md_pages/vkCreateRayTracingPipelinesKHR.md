# vkCreateRayTracingPipelinesKHR(3) Manual Page

## Name

vkCreateRayTracingPipelinesKHR - Creates a new ray tracing pipeline
object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create ray tracing pipelines, call:

``` c
// Provided by VK_KHR_ray_tracing_pipeline
VkResult vkCreateRayTracingPipelinesKHR(
    VkDevice                                    device,
    VkDeferredOperationKHR                      deferredOperation,
    VkPipelineCache                             pipelineCache,
    uint32_t                                    createInfoCount,
    const VkRayTracingPipelineCreateInfoKHR*    pCreateInfos,
    const VkAllocationCallbacks*                pAllocator,
    VkPipeline*                                 pPipelines);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the ray tracing pipelines.

- `deferredOperation` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or the
  handle of a valid
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#deferred-host-operations-requesting"
  target="_blank" rel="noopener">request deferral</a> object for this
  command.

- `pipelineCache` is either [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  indicating that pipeline caching is disabled, or the handle of a valid
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-cache"
  target="_blank" rel="noopener">pipeline cache</a> object, in which
  case use of that cache is enabled for the duration of the command.

- `createInfoCount` is the length of the `pCreateInfos` and `pPipelines`
  arrays.

- `pCreateInfos` is a pointer to an array of
  [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html)
  structures.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pPipelines` is a pointer to an array in which the resulting ray
  tracing pipeline objects are returned.

## <a href="#_description" class="anchor"></a>Description

The `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS` error is returned if the
implementation is unable to reuse the shader group handles provided in
[VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingShaderGroupCreateInfoKHR.html)::`pShaderGroupCaptureReplayHandle`
when
[VkPhysicalDeviceRayTracingPipelineFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelineFeaturesKHR.html)::`rayTracingPipelineShaderGroupHandleCaptureReplay`
is enabled.

Pipelines are created and returned as described for <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-multiple"
target="_blank" rel="noopener">Multiple Pipeline Creation</a>.

Valid Usage

- <a href="#VUID-vkCreateRayTracingPipelinesKHR-device-09677"
  id="VUID-vkCreateRayTracingPipelinesKHR-device-09677"></a>
  VUID-vkCreateRayTracingPipelinesKHR-device-09677  
  `device` **must** support at least one queue family with the
  `VK_QUEUE_COMPUTE_BIT` capability

- <a href="#VUID-vkCreateRayTracingPipelinesKHR-flags-03415"
  id="VUID-vkCreateRayTracingPipelinesKHR-flags-03415"></a>
  VUID-vkCreateRayTracingPipelinesKHR-flags-03415  
  If the `flags` member of any element of `pCreateInfos` contains the
  `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, and the `basePipelineIndex`
  member of that same element is not `-1`, `basePipelineIndex` **must**
  be less than the index into `pCreateInfos` that corresponds to that
  element

- <a href="#VUID-vkCreateRayTracingPipelinesKHR-flags-03416"
  id="VUID-vkCreateRayTracingPipelinesKHR-flags-03416"></a>
  VUID-vkCreateRayTracingPipelinesKHR-flags-03416  
  If the `flags` member of any element of `pCreateInfos` contains the
  `VK_PIPELINE_CREATE_DERIVATIVE_BIT` flag, the base pipeline **must**
  have been created with the `VK_PIPELINE_CREATE_ALLOW_DERIVATIVES_BIT`
  flag set

- <a href="#VUID-vkCreateRayTracingPipelinesKHR-flags-03816"
  id="VUID-vkCreateRayTracingPipelinesKHR-flags-03816"></a>
  VUID-vkCreateRayTracingPipelinesKHR-flags-03816  
  `flags` **must** not contain the `VK_PIPELINE_CREATE_DISPATCH_BASE`
  flag

- <a href="#VUID-vkCreateRayTracingPipelinesKHR-pipelineCache-02903"
  id="VUID-vkCreateRayTracingPipelinesKHR-pipelineCache-02903"></a>
  VUID-vkCreateRayTracingPipelinesKHR-pipelineCache-02903  
  If `pipelineCache` was created with
  `VK_PIPELINE_CACHE_CREATE_EXTERNALLY_SYNCHRONIZED_BIT`, host access to
  `pipelineCache` **must** be [externally
  synchronized](#fundamentals-threadingbehavior)

<!-- -->

- <a href="#VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-03678"
  id="VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-03678"></a>
  VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-03678  
  Any previous deferred operation that was associated with
  `deferredOperation` **must** be complete

- <a href="#VUID-vkCreateRayTracingPipelinesKHR-rayTracingPipeline-03586"
  id="VUID-vkCreateRayTracingPipelinesKHR-rayTracingPipeline-03586"></a>
  VUID-vkCreateRayTracingPipelinesKHR-rayTracingPipeline-03586  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-rayTracingPipeline"
  target="_blank" rel="noopener"><code>rayTracingPipeline</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-03587"
  id="VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-03587"></a>
  VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-03587  
  If `deferredOperation` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  the `flags` member of elements of `pCreateInfos` **must** not include
  `VK_PIPELINE_CREATE_EARLY_RETURN_ON_FAILURE_BIT`

Valid Usage (Implicit)

- <a href="#VUID-vkCreateRayTracingPipelinesKHR-device-parameter"
  id="VUID-vkCreateRayTracingPipelinesKHR-device-parameter"></a>
  VUID-vkCreateRayTracingPipelinesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-parameter"
  id="VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-parameter"></a>
  VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-parameter  
  If `deferredOperation` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `deferredOperation` **must** be a valid
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) handle

- <a href="#VUID-vkCreateRayTracingPipelinesKHR-pipelineCache-parameter"
  id="VUID-vkCreateRayTracingPipelinesKHR-pipelineCache-parameter"></a>
  VUID-vkCreateRayTracingPipelinesKHR-pipelineCache-parameter  
  If `pipelineCache` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `pipelineCache` **must** be a valid
  [VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html) handle

- <a href="#VUID-vkCreateRayTracingPipelinesKHR-pCreateInfos-parameter"
  id="VUID-vkCreateRayTracingPipelinesKHR-pCreateInfos-parameter"></a>
  VUID-vkCreateRayTracingPipelinesKHR-pCreateInfos-parameter  
  `pCreateInfos` **must** be a valid pointer to an array of
  `createInfoCount` valid
  [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html)
  structures

- <a href="#VUID-vkCreateRayTracingPipelinesKHR-pAllocator-parameter"
  id="VUID-vkCreateRayTracingPipelinesKHR-pAllocator-parameter"></a>
  VUID-vkCreateRayTracingPipelinesKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateRayTracingPipelinesKHR-pPipelines-parameter"
  id="VUID-vkCreateRayTracingPipelinesKHR-pPipelines-parameter"></a>
  VUID-vkCreateRayTracingPipelinesKHR-pPipelines-parameter  
  `pPipelines` **must** be a valid pointer to an array of
  `createInfoCount` [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handles

- <a
  href="#VUID-vkCreateRayTracingPipelinesKHR-createInfoCount-arraylength"
  id="VUID-vkCreateRayTracingPipelinesKHR-createInfoCount-arraylength"></a>
  VUID-vkCreateRayTracingPipelinesKHR-createInfoCount-arraylength  
  `createInfoCount` **must** be greater than `0`

- <a href="#VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-parent"
  id="VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-parent"></a>
  VUID-vkCreateRayTracingPipelinesKHR-deferredOperation-parent  
  If `deferredOperation` is a valid handle, it **must** have been
  created, allocated, or retrieved from `device`

- <a href="#VUID-vkCreateRayTracingPipelinesKHR-pipelineCache-parent"
  id="VUID-vkCreateRayTracingPipelinesKHR-pipelineCache-parent"></a>
  VUID-vkCreateRayTracingPipelinesKHR-pipelineCache-parent  
  If `pipelineCache` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_OPERATION_DEFERRED_KHR`

- `VK_OPERATION_NOT_DEFERRED_KHR`

- `VK_PIPELINE_COMPILE_REQUIRED_EXT`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCache.html),
[VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateRayTracingPipelinesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
