# VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE(3) Manual Page

## Name

VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE - Declare a non-dispatchable object handle



## [](#_c_specification)C Specification

`VK_DEFINE_NON_DISPATCHABLE_HANDLE` defines a [non-dispatchable handle](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-objectmodel-overview) type.

```c++
// Provided by VK_VERSION_1_0
#ifndef VK_DEFINE_NON_DISPATCHABLE_HANDLE
    #if (VK_USE_64_BIT_PTR_DEFINES==1)
        #define VK_DEFINE_NON_DISPATCHABLE_HANDLE(object) typedef struct object##_T *object;
    #else
        #define VK_DEFINE_NON_DISPATCHABLE_HANDLE(object) typedef uint64_t object;
    #endif
#endif
```

## [](#_description)Description

- `object` is the name of the resulting C type.

Most Vulkan handle types, such as [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), are non-dispatchable.

Note

The `vulkan_core.h` header allows the [VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html) and [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) definitions to be overridden by the application. If [VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html) is already defined when `vulkan_core.h` is compiled, the default definitions for [VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html) and [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) are skipped. This allows the application to define a binary-compatible custom handle which **may** provide more type-safety or other features needed by the application. Applications **must** not define handles in a way that is not binary compatible - where binary compatibility is platform dependent.

## [](#_see_also)See Also

[VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html), [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html), [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html), [VkCommandPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandPool.html), [VkCuFunctionNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuFunctionNVX.html), [VkCuModuleNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCuModuleNVX.html), [VkCudaFunctionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaFunctionNV.html), [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaModuleNV.html), [VkDataGraphPipelineSessionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionARM.html), [VkDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportCallbackEXT.html), [VkDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugUtilsMessengerEXT.html), [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeferredOperationKHR.html), [VkDescriptorPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPool.html), [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html), [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayout.html), [VkDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplate.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceMemory.html), [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html), [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeKHR.html), [VkEvent](https://registry.khronos.org/vulkan/specs/latest/man/html/VkEvent.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html), [VkFramebuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebuffer.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkImageView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageView.html), [VkIndirectCommandsLayoutEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutEXT.html), [VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectCommandsLayoutNV.html), [VkIndirectExecutionSetEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndirectExecutionSetEXT.html), [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html), [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionNV.html), [VkPerformanceConfigurationINTEL](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPerformanceConfigurationINTEL.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkPipelineBinaryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKHR.html), [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html), [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineLayout.html), [VkPrivateDataSlot](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrivateDataSlot.html), [VkQueryPool](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPool.html), [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html), [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html), [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversion.html), [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html), [VkShaderEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderEXT.html), [VkShaderModule](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderModule.html), [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceKHR.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html), [VkTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorARM.html), [VkTensorViewARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorViewARM.html), [VkValidationCacheEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCacheEXT.html), [VkVideoSessionKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionKHR.html), [VkVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVideoSessionParametersKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_DEFINE_NON_DISPATCHABLE_HANDLE).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0