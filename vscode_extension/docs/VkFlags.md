# VkFlags(3) Manual Page

## Name

VkFlags - Vulkan bitmasks



## <a href="#_c_specification" class="anchor"></a>C Specification

A collection of flags is represented by a bitmask using the type
[VkFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFlags.html):

``` c
// Provided by VK_VERSION_1_0
typedef uint32_t VkFlags;
```

## <a href="#_description" class="anchor"></a>Description

Bitmasks are passed to many commands and structures to compactly
represent options, but [VkFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFlags.html) is not used directly in
the API. Instead, a `Vk*Flags` type which is an alias of
[VkFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFlags.html), and whose name matches the corresponding
`Vk*FlagBits` that are valid for that type, is used.

Any `Vk*Flags` member or parameter used in the API as an input **must**
be a valid combination of bit flags. A valid combination is either zero
or the bitwise OR of valid bit flags.

An individual bit flag is valid for a `Vk*Flags` type if it would be a
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fundamentals-validusage-enums"
target="_blank" rel="noopener">valid enumerant</a> when used with the
equivalent `Vk*FlagBits` type, where the bits type is obtained by taking
the flag type and replacing the trailing `Flags` with `FlagBits`. For
example, a flag value of type
[VkColorComponentFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlags.html) **must** contain
only bit flags defined by
[VkColorComponentFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlagBits.html).

Any `Vk*Flags` member or parameter returned from a query command or
otherwise output from Vulkan to the application **may** contain bit
flags undefined in its corresponding `Vk*FlagBits` type. An application
**cannot** rely on the state of these unspecified bits.

Only the low-order 31 bits (bit positions zero through 30) are available
for use as flag bits.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>This restriction is due to poorly defined behavior by C compilers
given a C enumerant value of <code>0x80000000</code>. In some cases
adding this enumerant value may increase the size of the underlying
<code>Vk*FlagBits</code> type, breaking the ABI.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAccelerationStructureCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateFlagsKHR.html),
[VkAccelerationStructureMotionInfoFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInfoFlagsNV.html),
[VkAccelerationStructureMotionInstanceFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureMotionInstanceFlagsNV.html),
[VkAccessFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlags.html),
[VkAcquireProfilingLockFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAcquireProfilingLockFlagsKHR.html),
[VkAndroidSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidSurfaceCreateFlagsKHR.html),
[VkAttachmentDescriptionFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescriptionFlags.html),
[VkBufferCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateFlags.html),
[VkBufferUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferUsageFlags.html),
[VkBufferViewCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferViewCreateFlags.html),
[VkBuildAccelerationStructureFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildAccelerationStructureFlagsKHR.html),
[VkBuildMicromapFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildMicromapFlagsEXT.html),
[VkColorComponentFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkColorComponentFlags.html),
[VkCommandBufferResetFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferResetFlags.html),
[VkCommandBufferUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferUsageFlags.html),
[VkCommandPoolCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolCreateFlags.html),
[VkCommandPoolResetFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolResetFlags.html),
[VkCommandPoolTrimFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolTrimFlags.html),
[VkCompositeAlphaFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompositeAlphaFlagsKHR.html),
[VkConditionalRenderingFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkConditionalRenderingFlagsEXT.html),
[VkCullModeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCullModeFlags.html),
[VkDebugReportFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportFlagsEXT.html),
[VkDebugUtilsMessageSeverityFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageSeverityFlagsEXT.html),
[VkDebugUtilsMessageTypeFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessageTypeFlagsEXT.html),
[VkDebugUtilsMessengerCallbackDataFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCallbackDataFlagsEXT.html),
[VkDebugUtilsMessengerCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugUtilsMessengerCreateFlagsEXT.html),
[VkDependencyFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDependencyFlags.html),
[VkDescriptorBindingFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlags.html),
[VkDescriptorPoolCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateFlags.html),
[VkDescriptorPoolResetFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolResetFlags.html),
[VkDescriptorSetLayoutCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateFlags.html),
[VkDescriptorUpdateTemplateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateCreateFlags.html),
[VkDeviceAddressBindingFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddressBindingFlagsEXT.html),
[VkDeviceCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateFlags.html),
[VkDeviceDiagnosticsConfigFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceDiagnosticsConfigFlagsNV.html),
[VkDeviceGroupPresentModeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagsKHR.html),
[VkDeviceMemoryReportFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemoryReportFlagsEXT.html),
[VkDeviceQueueCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceQueueCreateFlags.html),
[VkDirectDriverLoadingFlagsLUNARG](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectDriverLoadingFlagsLUNARG.html),
[VkDirectFBSurfaceCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDirectFBSurfaceCreateFlagsEXT.html),
[VkDisplayModeCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeCreateFlagsKHR.html),
[VkDisplayPlaneAlphaFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneAlphaFlagsKHR.html),
[VkDisplaySurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplaySurfaceCreateFlagsKHR.html),
[VkEventCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkEventCreateFlags.html),
[VkExportMetalObjectTypeFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectTypeFlagsEXT.html),
[VkExternalFenceFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceFeatureFlags.html),
[VkExternalFenceHandleTypeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFenceHandleTypeFlags.html),
[VkExternalMemoryFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlags.html),
[VkExternalMemoryFeatureFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlagsNV.html),
[VkExternalMemoryHandleTypeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlags.html),
[VkExternalMemoryHandleTypeFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagsNV.html),
[VkExternalSemaphoreFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreFeatureFlags.html),
[VkExternalSemaphoreHandleTypeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalSemaphoreHandleTypeFlags.html),
[VkFenceCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceCreateFlags.html),
[VkFenceImportFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFenceImportFlags.html),
[VkFlags64](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFlags64.html),
[VkFormatFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlags.html),
[VkFrameBoundaryFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFrameBoundaryFlagsEXT.html),
[VkFramebufferCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferCreateFlags.html),
[VkGeometryFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryFlagsKHR.html),
[VkGeometryInstanceFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryInstanceFlagsKHR.html),
[VkGraphicsPipelineLibraryFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineLibraryFlagsEXT.html),
[VkHeadlessSurfaceCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHeadlessSurfaceCreateFlagsEXT.html),
[VkHostImageCopyFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkHostImageCopyFlagsEXT.html),
[VkIOSSurfaceCreateFlagsMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIOSSurfaceCreateFlagsMVK.html),
[VkImageAspectFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageAspectFlags.html),
[VkImageCompressionFixedRateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionFixedRateFlagsEXT.html),
[VkImageCompressionFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCompressionFlagsEXT.html),
[VkImageConstraintsInfoFlagsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageConstraintsInfoFlagsFUCHSIA.html),
[VkImageCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlags.html),
[VkImageFormatConstraintsFlagsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatConstraintsFlagsFUCHSIA.html),
[VkImagePipeSurfaceCreateFlagsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePipeSurfaceCreateFlagsFUCHSIA.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkImageViewCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewCreateFlags.html),
[VkIndirectCommandsLayoutUsageFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutUsageFlagsNV.html),
[VkIndirectStateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectStateFlagsNV.html),
[VkInstanceCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateFlags.html),
[VkMacOSSurfaceCreateFlagsMVK](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMacOSSurfaceCreateFlagsMVK.html),
[VkMemoryAllocateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlags.html),
[VkMemoryHeapFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryHeapFlags.html),
[VkMemoryMapFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryMapFlags.html),
[VkMemoryPropertyFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryPropertyFlags.html),
[VkMemoryUnmapFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryUnmapFlagsKHR.html),
[VkMetalSurfaceCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMetalSurfaceCreateFlagsEXT.html),
[VkMicromapCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateFlagsEXT.html),
[VkOpticalFlowExecuteFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowExecuteFlagsNV.html),
[VkOpticalFlowGridSizeFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowGridSizeFlagsNV.html),
[VkOpticalFlowSessionCreateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionCreateFlagsNV.html),
[VkOpticalFlowUsageFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowUsageFlagsNV.html),
[VkPeerMemoryFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPeerMemoryFeatureFlags.html),
[VkPerformanceCounterDescriptionFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPerformanceCounterDescriptionFlagsKHR.html),
[VkPipelineCacheCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCacheCreateFlags.html),
[VkPipelineColorBlendStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateFlags.html),
[VkPipelineCompilerControlFlagsAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCompilerControlFlagsAMD.html),
[VkPipelineCoverageModulationStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageModulationStateCreateFlagsNV.html),
[VkPipelineCoverageReductionStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageReductionStateCreateFlagsNV.html),
[VkPipelineCoverageToColorStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageToColorStateCreateFlagsNV.html),
[VkPipelineCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreateFlags.html),
[VkPipelineCreationFeedbackFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackFlags.html),
[VkPipelineDepthStencilStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateFlags.html),
[VkPipelineDiscardRectangleStateCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDiscardRectangleStateCreateFlagsEXT.html),
[VkPipelineDynamicStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateFlags.html),
[VkPipelineInputAssemblyStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineInputAssemblyStateCreateFlags.html),
[VkPipelineLayoutCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateFlags.html),
[VkPipelineMultisampleStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateFlags.html),
[VkPipelineRasterizationConservativeStateCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationConservativeStateCreateFlagsEXT.html),
[VkPipelineRasterizationDepthClipStateCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationDepthClipStateCreateFlagsEXT.html),
[VkPipelineRasterizationStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateFlags.html),
[VkPipelineRasterizationStateStreamCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateStreamCreateFlagsEXT.html),
[VkPipelineShaderStageCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateFlags.html),
[VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags.html),
[VkPipelineTessellationStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineTessellationStateCreateFlags.html),
[VkPipelineVertexInputStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputStateCreateFlags.html),
[VkPipelineViewportStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateFlags.html),
[VkPipelineViewportSwizzleStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportSwizzleStateCreateFlagsNV.html),
[VkPresentGravityFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentGravityFlagsEXT.html),
[VkPresentScalingFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentScalingFlagsEXT.html),
[VkPrivateDataSlotCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPrivateDataSlotCreateFlags.html),
[VkQueryControlFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryControlFlags.html),
[VkQueryPipelineStatisticFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlags.html),
[VkQueryPoolCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPoolCreateFlags.html),
[VkQueryResultFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryResultFlags.html),
[VkQueueFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFlags.html),
[VkRenderPassCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateFlags.html),
[VkRenderingFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingFlags.html),
[VkResolveModeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveModeFlags.html),
[VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlags.html),
[VkSamplerCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateFlags.html),
[VkScreenSurfaceCreateFlagsQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenSurfaceCreateFlagsQNX.html),
[VkSemaphoreCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateFlags.html),
[VkSemaphoreImportFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlags.html),
[VkSemaphoreWaitFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreWaitFlags.html),
[VkShaderCorePropertiesFlagsAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCorePropertiesFlagsAMD.html),
[VkShaderCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateFlagsEXT.html),
[VkShaderModuleCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderModuleCreateFlags.html),
[VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlags.html),
[VkSparseImageFormatFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseImageFormatFlags.html),
[VkSparseMemoryBindFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSparseMemoryBindFlags.html),
[VkStencilFaceFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilFaceFlags.html),
[VkStreamDescriptorSurfaceCreateFlagsGGP](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStreamDescriptorSurfaceCreateFlagsGGP.html),
[VkSubgroupFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubgroupFeatureFlags.html),
[VkSubmitFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubmitFlags.html),
[VkSubpassDescriptionFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescriptionFlags.html),
[VkSurfaceCounterFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCounterFlagsEXT.html),
[VkSurfaceTransformFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagsKHR.html),
[VkSwapchainCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateFlagsKHR.html),
[VkToolPurposeFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkToolPurposeFlags.html),
[VkValidationCacheCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCacheCreateFlagsEXT.html),
[VkViSurfaceCreateFlagsNN](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViSurfaceCreateFlagsNN.html),
[VkVideoBeginCodingFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoBeginCodingFlagsKHR.html),
[VkVideoCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCapabilityFlagsKHR.html),
[VkVideoChromaSubsamplingFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoChromaSubsamplingFlagsKHR.html),
[VkVideoCodecOperationFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodecOperationFlagsKHR.html),
[VkVideoCodingControlFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoCodingControlFlagsKHR.html),
[VkVideoComponentBitDepthFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoComponentBitDepthFlagsKHR.html),
[VkVideoDecodeCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeCapabilityFlagsKHR.html),
[VkVideoDecodeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeFlagsKHR.html),
[VkVideoDecodeH264PictureLayoutFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeH264PictureLayoutFlagsKHR.html),
[VkVideoDecodeUsageFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoDecodeUsageFlagsKHR.html),
[VkVideoEncodeCapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeCapabilityFlagsKHR.html),
[VkVideoEncodeContentFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeContentFlagsKHR.html),
[VkVideoEncodeFeedbackFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeFeedbackFlagsKHR.html),
[VkVideoEncodeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeFlagsKHR.html),
[VkVideoEncodeH264CapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264CapabilityFlagsKHR.html),
[VkVideoEncodeH264RateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264RateControlFlagsKHR.html),
[VkVideoEncodeH264StdFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH264StdFlagsKHR.html),
[VkVideoEncodeH265CapabilityFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CapabilityFlagsKHR.html),
[VkVideoEncodeH265CtbSizeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265CtbSizeFlagsKHR.html),
[VkVideoEncodeH265RateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265RateControlFlagsKHR.html),
[VkVideoEncodeH265StdFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265StdFlagsKHR.html),
[VkVideoEncodeH265TransformBlockSizeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeH265TransformBlockSizeFlagsKHR.html),
[VkVideoEncodeRateControlFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlFlagsKHR.html),
[VkVideoEncodeRateControlModeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeRateControlModeFlagsKHR.html),
[VkVideoEncodeUsageFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEncodeUsageFlagsKHR.html),
[VkVideoEndCodingFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoEndCodingFlagsKHR.html),
[VkVideoSessionCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionCreateFlagsKHR.html),
[VkVideoSessionParametersCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVideoSessionParametersCreateFlagsKHR.html),
[VkWaylandSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWaylandSurfaceCreateFlagsKHR.html),
[VkWin32SurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWin32SurfaceCreateFlagsKHR.html),
[VkXcbSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkXcbSurfaceCreateFlagsKHR.html),
[VkXlibSurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkXlibSurfaceCreateFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkFlags"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
