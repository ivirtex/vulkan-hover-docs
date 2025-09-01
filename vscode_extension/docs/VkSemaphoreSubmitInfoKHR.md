# VkSemaphoreSubmitInfo(3) Manual Page

## Name

VkSemaphoreSubmitInfo - Structure specifying a semaphore signal or wait operation



## [](#_c_specification)C Specification

The `VkSemaphoreSubmitInfo` structure is defined as:

```c++
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

```c++
// Provided by VK_KHR_synchronization2
typedef VkSemaphoreSubmitInfo VkSemaphoreSubmitInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `semaphore` is a [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) affected by this operation.
- `value` is either the value used to signal `semaphore` or the value waited on by `semaphore`, if `semaphore` is a timeline semaphore. Otherwise it is ignored.
- `stageMask` is a [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html) mask of pipeline stages which limit the first synchronization scope of a semaphore signal operation, or second synchronization scope of a semaphore wait operation as described in the [semaphore wait operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-waiting) and [semaphore signal operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-semaphores-signaling) sections of [the synchronization chapter](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization).
- `deviceIndex` is the index of the device within a device group that executes the semaphore wait or signal operation.

## [](#_description)Description

Whether this structure defines a semaphore wait or signal operation is defined by how it is used.

Valid Usage

- [](#VUID-VkSemaphoreSubmitInfo-stageMask-03929)VUID-VkSemaphoreSubmitInfo-stageMask-03929  
  If the [`geometryShader`](#features-geometryShader) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_GEOMETRY_SHADER_BIT`
- [](#VUID-VkSemaphoreSubmitInfo-stageMask-03930)VUID-VkSemaphoreSubmitInfo-stageMask-03930  
  If the [`tessellationShader`](#features-tessellationShader) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_TESSELLATION_CONTROL_SHADER_BIT` or `VK_PIPELINE_STAGE_2_TESSELLATION_EVALUATION_SHADER_BIT`
- [](#VUID-VkSemaphoreSubmitInfo-stageMask-03931)VUID-VkSemaphoreSubmitInfo-stageMask-03931  
  If the [`conditionalRendering`](#features-conditionalRendering) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_CONDITIONAL_RENDERING_BIT_EXT`
- [](#VUID-VkSemaphoreSubmitInfo-stageMask-03932)VUID-VkSemaphoreSubmitInfo-stageMask-03932  
  If the [`fragmentDensityMap`](#features-fragmentDensityMap) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_FRAGMENT_DENSITY_PROCESS_BIT_EXT`
- [](#VUID-VkSemaphoreSubmitInfo-stageMask-03933)VUID-VkSemaphoreSubmitInfo-stageMask-03933  
  If the [`transformFeedback`](#features-transformFeedback) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_TRANSFORM_FEEDBACK_BIT_EXT`
- [](#VUID-VkSemaphoreSubmitInfo-stageMask-03934)VUID-VkSemaphoreSubmitInfo-stageMask-03934  
  If the [`meshShader`](#features-meshShader) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_MESH_SHADER_BIT_EXT`
- [](#VUID-VkSemaphoreSubmitInfo-stageMask-03935)VUID-VkSemaphoreSubmitInfo-stageMask-03935  
  If the [`taskShader`](#features-taskShader) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_TASK_SHADER_BIT_EXT`
- [](#VUID-VkSemaphoreSubmitInfo-stageMask-07316)VUID-VkSemaphoreSubmitInfo-stageMask-07316  
  If neither of the [`shadingRateImage`](#features-shadingRateImage) or the [`attachmentFragmentShadingRate`](#features-attachmentFragmentShadingRate) features are enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_FRAGMENT_SHADING_RATE_ATTACHMENT_BIT_KHR`
- [](#VUID-VkSemaphoreSubmitInfo-stageMask-04957)VUID-VkSemaphoreSubmitInfo-stageMask-04957  
  If the [`subpassShading`](#features-subpassShading) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_SUBPASS_SHADER_BIT_HUAWEI`
- [](#VUID-VkSemaphoreSubmitInfo-stageMask-04995)VUID-VkSemaphoreSubmitInfo-stageMask-04995  
  If the [`invocationMask`](#features-invocationMask) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`
- [](#VUID-VkSemaphoreSubmitInfo-stageMask-07946)VUID-VkSemaphoreSubmitInfo-stageMask-07946  
  If neither the [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html) extension or the [`rayTracingPipeline`](#features-rayTracingPipeline) feature are enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_RAY_TRACING_SHADER_BIT_KHR`
- [](#VUID-VkSemaphoreSubmitInfo-stageMask-10751)VUID-VkSemaphoreSubmitInfo-stageMask-10751  
  If the [`accelerationStructure`](#features-accelerationStructure) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_BUILD_BIT_KHR`
- [](#VUID-VkSemaphoreSubmitInfo-stageMask-10752)VUID-VkSemaphoreSubmitInfo-stageMask-10752  
  If the [`rayTracingMaintenance1`](#features-rayTracingMaintenance1) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR`
- [](#VUID-VkSemaphoreSubmitInfo-stageMask-10753)VUID-VkSemaphoreSubmitInfo-stageMask-10753  
  If the [`micromap`](#features-micromap) feature is not enabled, `stageMask` **must** not contain `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT`
- [](#VUID-VkSemaphoreSubmitInfo-device-03888)VUID-VkSemaphoreSubmitInfo-device-03888  
  If the `device` that `semaphore` was created on is not a device group, `deviceIndex` **must** be `0`
- [](#VUID-VkSemaphoreSubmitInfo-device-03889)VUID-VkSemaphoreSubmitInfo-device-03889  
  If the `device` that `semaphore` was created on is a device group, `deviceIndex` **must** be a valid device index

Valid Usage (Implicit)

- [](#VUID-VkSemaphoreSubmitInfo-sType-sType)VUID-VkSemaphoreSubmitInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SEMAPHORE_SUBMIT_INFO`
- [](#VUID-VkSemaphoreSubmitInfo-pNext-pNext)VUID-VkSemaphoreSubmitInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkSemaphoreSubmitInfo-semaphore-parameter)VUID-VkSemaphoreSubmitInfo-semaphore-parameter  
  `semaphore` **must** be a valid [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) handle
- [](#VUID-VkSemaphoreSubmitInfo-stageMask-parameter)VUID-VkSemaphoreSubmitInfo-stageMask-parameter  
  `stageMask` **must** be a valid combination of [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html) values

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html), [VkRenderPassStripeSubmitInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassStripeSubmitInfoARM.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSubmitInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubmitInfo2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSemaphoreSubmitInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0