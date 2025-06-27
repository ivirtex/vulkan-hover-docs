# VkSemaphoreSubmitInfo(3) Manual Page

## Name

VkSemaphoreSubmitInfo - Structure specifying a semaphore signal or wait
operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSemaphoreSubmitInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkSemaphoreSubmitInfo {
    VkStructureType          sType;
    const void*              pNext;
    VkSemaphore              semaphore;
    uint64_t                 value;
    VkPipelineStageFlags2    stageMask;
    uint32_t                 deviceIndex;
} VkSemaphoreSubmitInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_synchronization2
typedef VkSemaphoreSubmitInfo VkSemaphoreSubmitInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `semaphore` is a [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) affected by this
  operation.

- `value` is either the value used to signal `semaphore` or the value
  waited on by `semaphore`, if `semaphore` is a timeline semaphore.
  Otherwise it is ignored.

- `stageMask` is a [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags2.html)
  mask of pipeline stages which limit the first synchronization scope of
  a semaphore signal operation, or second synchronization scope of a
  semaphore wait operation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-waiting"
  target="_blank" rel="noopener">semaphore wait operation</a> and <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-semaphores-signaling"
  target="_blank" rel="noopener">semaphore signal operation</a> sections
  of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization"
  target="_blank" rel="noopener">the synchronization chapter</a>.

- `deviceIndex` is the index of the device within a device group that
  executes the semaphore wait or signal operation.

## <a href="#_description" class="anchor"></a>Description

Whether this structure defines a semaphore wait or signal operation is
defined by how it is used.

Valid Usage

- <a href="#VUID-VkSemaphoreSubmitInfo-stageMask-03929"
  id="VUID-VkSemaphoreSubmitInfo-stageMask-03929"></a>
  VUID-VkSemaphoreSubmitInfo-stageMask-03929  
  If the [`geometryShader`](#features-geometryShader) feature is not
  enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`

- <a href="#VUID-VkSemaphoreSubmitInfo-stageMask-03930"
  id="VUID-VkSemaphoreSubmitInfo-stageMask-03930"></a>
  VUID-VkSemaphoreSubmitInfo-stageMask-03930  
  If the [`tessellationShader`](#features-tessellationShader) feature is
  not enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT` or
  `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`

- <a href="#VUID-VkSemaphoreSubmitInfo-stageMask-03931"
  id="VUID-VkSemaphoreSubmitInfo-stageMask-03931"></a>
  VUID-VkSemaphoreSubmitInfo-stageMask-03931  
  If the [`conditionalRendering`](#features-conditionalRendering)
  feature is not enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`

- <a href="#VUID-VkSemaphoreSubmitInfo-stageMask-03932"
  id="VUID-VkSemaphoreSubmitInfo-stageMask-03932"></a>
  VUID-VkSemaphoreSubmitInfo-stageMask-03932  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is
  not enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`

- <a href="#VUID-VkSemaphoreSubmitInfo-stageMask-03933"
  id="VUID-VkSemaphoreSubmitInfo-stageMask-03933"></a>
  VUID-VkSemaphoreSubmitInfo-stageMask-03933  
  If the [`transformFeedback`](#features-transformFeedback) feature is
  not enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`

- <a href="#VUID-VkSemaphoreSubmitInfo-stageMask-03934"
  id="VUID-VkSemaphoreSubmitInfo-stageMask-03934"></a>
  VUID-VkSemaphoreSubmitInfo-stageMask-03934  
  If the [`meshShader`](#features-meshShader) feature is not enabled,
  `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`

- <a href="#VUID-VkSemaphoreSubmitInfo-stageMask-03935"
  id="VUID-VkSemaphoreSubmitInfo-stageMask-03935"></a>
  VUID-VkSemaphoreSubmitInfo-stageMask-03935  
  If the [`taskShader`](#features-taskShader) feature is not enabled,
  `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`

- <a href="#VUID-VkSemaphoreSubmitInfo-stageMask-07316"
  id="VUID-VkSemaphoreSubmitInfo-stageMask-07316"></a>
  VUID-VkSemaphoreSubmitInfo-stageMask-07316  
  If neither the [`shadingRateImage`](#features-shadingRateImage) or
  [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate)
  are enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`

- <a href="#VUID-VkSemaphoreSubmitInfo-stageMask-04957"
  id="VUID-VkSemaphoreSubmitInfo-stageMask-04957"></a>
  VUID-VkSemaphoreSubmitInfo-stageMask-04957  
  If the [`subpassShading`](#features-subpassShading) feature is not
  enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`

- <a href="#VUID-VkSemaphoreSubmitInfo-stageMask-04995"
  id="VUID-VkSemaphoreSubmitInfo-stageMask-04995"></a>
  VUID-VkSemaphoreSubmitInfo-stageMask-04995  
  If the [`invocationMask`](#features-invocationMask) feature is not
  enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`

- <a href="#VUID-VkSemaphoreSubmitInfo-stageMask-07946"
  id="VUID-VkSemaphoreSubmitInfo-stageMask-07946"></a>
  VUID-VkSemaphoreSubmitInfo-stageMask-07946  
  If neither the [VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html) extension
  or [`rayTracingPipeline` feature](#features-rayTracingPipeline) are
  enabled, `stageMask` **must** not contain
  `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`

- <a href="#VUID-VkSemaphoreSubmitInfo-device-03888"
  id="VUID-VkSemaphoreSubmitInfo-device-03888"></a>
  VUID-VkSemaphoreSubmitInfo-device-03888  
  If the `device` that `semaphore` was created on is not a device group,
  `deviceIndex` **must** be `0`

- <a href="#VUID-VkSemaphoreSubmitInfo-device-03889"
  id="VUID-VkSemaphoreSubmitInfo-device-03889"></a>
  VUID-VkSemaphoreSubmitInfo-device-03889  
  If the `device` that `semaphore` was created on is a device group,
  `deviceIndex` **must** be a valid device index

Valid Usage (Implicit)

- <a href="#VUID-VkSemaphoreSubmitInfo-sType-sType"
  id="VUID-VkSemaphoreSubmitInfo-sType-sType"></a>
  VUID-VkSemaphoreSubmitInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SEMAPHORE_SUBMIT_INFO`

- <a href="#VUID-VkSemaphoreSubmitInfo-pNext-pNext"
  id="VUID-VkSemaphoreSubmitInfo-pNext-pNext"></a>
  VUID-VkSemaphoreSubmitInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkSemaphoreSubmitInfo-semaphore-parameter"
  id="VUID-VkSemaphoreSubmitInfo-semaphore-parameter"></a>
  VUID-VkSemaphoreSubmitInfo-semaphore-parameter  
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html) handle

- <a href="#VUID-VkSemaphoreSubmitInfo-stageMask-parameter"
  id="VUID-VkSemaphoreSubmitInfo-stageMask-parameter"></a>
  VUID-VkSemaphoreSubmitInfo-stageMask-parameter  
  `stageMask` **must** be a valid combination of
  [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags2.html),
[VkRenderPassStripeSubmitInfoARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassStripeSubmitInfoARM.html),
[VkSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphore.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitInfo2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSemaphoreSubmitInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
