# VkAllocationCallbacks(3) Manual Page

## Name

VkAllocationCallbacks - Structure containing callback function pointers
for memory allocation



## <a href="#_c_specification" class="anchor"></a>C Specification

Allocators are provided by the application as a pointer to a
`VkAllocationCallbacks` structure:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkAllocationCallbacks {
    void*                                   pUserData;
    PFN_vkAllocationFunction                pfnAllocation;
    PFN_vkReallocationFunction              pfnReallocation;
    PFN_vkFreeFunction                      pfnFree;
    PFN_vkInternalAllocationNotification    pfnInternalAllocation;
    PFN_vkInternalFreeNotification          pfnInternalFree;
} VkAllocationCallbacks;
```

## <a href="#_members" class="anchor"></a>Members

## <a href="#_description" class="anchor"></a>Description

- `pUserData` is a value to be interpreted by the implementation of the
  callbacks. When any of the callbacks in `VkAllocationCallbacks` are
  called, the Vulkan implementation will pass this value as the first
  parameter to the callback. This value **can** vary each time an
  allocator is passed into a command, even when the same object takes an
  allocator in multiple commands.

- `pfnAllocation` is a
  [PFN_vkAllocationFunction](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkAllocationFunction.html) pointer to
  an application-defined memory allocation function.

- `pfnReallocation` is a
  [PFN_vkReallocationFunction](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkReallocationFunction.html) pointer
  to an application-defined memory reallocation function.

- `pfnFree` is a [PFN_vkFreeFunction](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkFreeFunction.html) pointer
  to an application-defined memory free function.

- `pfnInternalAllocation` is a
  [PFN_vkInternalAllocationNotification](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkInternalAllocationNotification.html)
  pointer to an application-defined function that is called by the
  implementation when the implementation makes internal allocations.

- `pfnInternalFree` is a
  [PFN_vkInternalFreeNotification](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkInternalFreeNotification.html)
  pointer to an application-defined function that is called by the
  implementation when the implementation frees internal allocations.

Valid Usage

- <a href="#VUID-VkAllocationCallbacks-pfnAllocation-00632"
  id="VUID-VkAllocationCallbacks-pfnAllocation-00632"></a>
  VUID-VkAllocationCallbacks-pfnAllocation-00632  
  `pfnAllocation` **must** be a valid pointer to a valid
  application-defined
  [PFN_vkAllocationFunction](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkAllocationFunction.html)

- <a href="#VUID-VkAllocationCallbacks-pfnReallocation-00633"
  id="VUID-VkAllocationCallbacks-pfnReallocation-00633"></a>
  VUID-VkAllocationCallbacks-pfnReallocation-00633  
  `pfnReallocation` **must** be a valid pointer to a valid
  application-defined
  [PFN_vkReallocationFunction](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkReallocationFunction.html)

- <a href="#VUID-VkAllocationCallbacks-pfnFree-00634"
  id="VUID-VkAllocationCallbacks-pfnFree-00634"></a>
  VUID-VkAllocationCallbacks-pfnFree-00634  
  `pfnFree` **must** be a valid pointer to a valid application-defined
  [PFN_vkFreeFunction](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkFreeFunction.html)

- <a href="#VUID-VkAllocationCallbacks-pfnInternalAllocation-00635"
  id="VUID-VkAllocationCallbacks-pfnInternalAllocation-00635"></a>
  VUID-VkAllocationCallbacks-pfnInternalAllocation-00635  
  If either of `pfnInternalAllocation` or `pfnInternalFree` is not
  `NULL`, both **must** be valid callbacks

## <a href="#_see_also" class="anchor"></a>See Also

[PFN_vkAllocationFunction](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkAllocationFunction.html),
[PFN_vkFreeFunction](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkFreeFunction.html),
[PFN_vkInternalAllocationNotification](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkInternalAllocationNotification.html),
[PFN_vkInternalFreeNotification](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkInternalFreeNotification.html),
[PFN_vkReallocationFunction](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/PFN_vkReallocationFunction.html),
[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[vkAllocateMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAllocateMemory.html),
[vkCreateAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateAccelerationStructureKHR.html),
[vkCreateAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateAccelerationStructureNV.html),
[vkCreateAndroidSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateAndroidSurfaceKHR.html),
[vkCreateBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateBuffer.html),
[vkCreateBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateBufferCollectionFUCHSIA.html),
[vkCreateBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateBufferView.html),
[vkCreateCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCommandPool.html),
[vkCreateComputePipelines](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateComputePipelines.html),
[vkCreateCuFunctionNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCuFunctionNVX.html),
[vkCreateCuModuleNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCuModuleNVX.html),
[vkCreateCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCudaFunctionNV.html),
[vkCreateCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCudaModuleNV.html),
[vkCreateDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDebugReportCallbackEXT.html),
[vkCreateDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDebugUtilsMessengerEXT.html),
[vkCreateDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDeferredOperationKHR.html),
[vkCreateDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorPool.html),
[vkCreateDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorSetLayout.html),
[vkCreateDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorUpdateTemplate.html),
[vkCreateDescriptorUpdateTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDescriptorUpdateTemplateKHR.html),
[vkCreateDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDevice.html),
[vkCreateDirectFBSurfaceEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDirectFBSurfaceEXT.html),
[vkCreateDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDisplayModeKHR.html),
[vkCreateDisplayPlaneSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDisplayPlaneSurfaceKHR.html),
[vkCreateEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateEvent.html),
[vkCreateExecutionGraphPipelinesAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateExecutionGraphPipelinesAMDX.html),
[vkCreateFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateFence.html),
[vkCreateFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateFramebuffer.html),
[vkCreateGraphicsPipelines](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateGraphicsPipelines.html),
[vkCreateHeadlessSurfaceEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateHeadlessSurfaceEXT.html),
[vkCreateIOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateIOSSurfaceMVK.html),
[vkCreateImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html),
[vkCreateImagePipeSurfaceFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImagePipeSurfaceFUCHSIA.html),
[vkCreateImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImageView.html),
[vkCreateIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateIndirectCommandsLayoutNV.html),
[vkCreateInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateInstance.html),
[vkCreateMacOSSurfaceMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateMacOSSurfaceMVK.html),
[vkCreateMetalSurfaceEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateMetalSurfaceEXT.html),
[vkCreateMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateMicromapEXT.html),
[vkCreateOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateOpticalFlowSessionNV.html),
[vkCreatePipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreatePipelineCache.html),
[vkCreatePipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreatePipelineLayout.html),
[vkCreatePrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreatePrivateDataSlot.html),
[vkCreatePrivateDataSlotEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreatePrivateDataSlotEXT.html),
[vkCreateQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateQueryPool.html),
[vkCreateRayTracingPipelinesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRayTracingPipelinesKHR.html),
[vkCreateRayTracingPipelinesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRayTracingPipelinesNV.html),
[vkCreateRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass.html),
[vkCreateRenderPass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass2.html),
[vkCreateRenderPass2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass2KHR.html),
[vkCreateSampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSampler.html),
[vkCreateSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSamplerYcbcrConversion.html),
[vkCreateSamplerYcbcrConversionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSamplerYcbcrConversionKHR.html),
[vkCreateScreenSurfaceQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateScreenSurfaceQNX.html),
[vkCreateSemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSemaphore.html),
[vkCreateShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateShaderModule.html),
[vkCreateShadersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateShadersEXT.html),
[vkCreateSharedSwapchainsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSharedSwapchainsKHR.html),
[vkCreateStreamDescriptorSurfaceGGP](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateStreamDescriptorSurfaceGGP.html),
[vkCreateSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateSwapchainKHR.html),
[vkCreateValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateValidationCacheEXT.html),
[vkCreateViSurfaceNN](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateViSurfaceNN.html),
[vkCreateVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateVideoSessionKHR.html),
[vkCreateVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateVideoSessionParametersKHR.html),
[vkCreateWaylandSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateWaylandSurfaceKHR.html),
[vkCreateWin32SurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateWin32SurfaceKHR.html),
[vkCreateXcbSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateXcbSurfaceKHR.html),
[vkCreateXlibSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateXlibSurfaceKHR.html),
[vkDestroyAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyAccelerationStructureKHR.html),
[vkDestroyAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyAccelerationStructureNV.html),
[vkDestroyBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyBuffer.html),
[vkDestroyBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyBufferCollectionFUCHSIA.html),
[vkDestroyBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyBufferView.html),
[vkDestroyCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyCommandPool.html),
[vkDestroyCuFunctionNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyCuFunctionNVX.html),
[vkDestroyCuModuleNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyCuModuleNVX.html),
[vkDestroyCudaFunctionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyCudaFunctionNV.html),
[vkDestroyCudaModuleNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyCudaModuleNV.html),
[vkDestroyDebugReportCallbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDebugReportCallbackEXT.html),
[vkDestroyDebugUtilsMessengerEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDebugUtilsMessengerEXT.html),
[vkDestroyDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDeferredOperationKHR.html),
[vkDestroyDescriptorPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorPool.html),
[vkDestroyDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorSetLayout.html),
[vkDestroyDescriptorUpdateTemplate](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorUpdateTemplate.html),
[vkDestroyDescriptorUpdateTemplateKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDescriptorUpdateTemplateKHR.html),
[vkDestroyDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyDevice.html),
[vkDestroyEvent](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyEvent.html),
[vkDestroyFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyFence.html),
[vkDestroyFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyFramebuffer.html),
[vkDestroyImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyImage.html),
[vkDestroyImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyImageView.html),
[vkDestroyIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyIndirectCommandsLayoutNV.html),
[vkDestroyInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyInstance.html),
[vkDestroyMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyMicromapEXT.html),
[vkDestroyOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyOpticalFlowSessionNV.html),
[vkDestroyPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyPipeline.html),
[vkDestroyPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyPipelineCache.html),
[vkDestroyPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyPipelineLayout.html),
[vkDestroyPrivateDataSlot](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyPrivateDataSlot.html),
[vkDestroyPrivateDataSlotEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyPrivateDataSlotEXT.html),
[vkDestroyQueryPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyQueryPool.html),
[vkDestroyRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyRenderPass.html),
[vkDestroySampler](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySampler.html),
[vkDestroySamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySamplerYcbcrConversion.html),
[vkDestroySamplerYcbcrConversionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySamplerYcbcrConversionKHR.html),
[vkDestroySemaphore](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySemaphore.html),
[vkDestroyShaderEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyShaderEXT.html),
[vkDestroyShaderModule](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyShaderModule.html),
[vkDestroySurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySurfaceKHR.html),
[vkDestroySwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroySwapchainKHR.html),
[vkDestroyValidationCacheEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyValidationCacheEXT.html),
[vkDestroyVideoSessionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyVideoSessionKHR.html),
[vkDestroyVideoSessionParametersKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyVideoSessionParametersKHR.html),
[vkFreeMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkFreeMemory.html),
[vkRegisterDeviceEventEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkRegisterDeviceEventEXT.html),
[vkRegisterDisplayEventEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkRegisterDisplayEventEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAllocationCallbacks"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
