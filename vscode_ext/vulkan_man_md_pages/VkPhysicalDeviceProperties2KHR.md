# VkPhysicalDeviceProperties2(3) Manual Page

## Name

VkPhysicalDeviceProperties2 - Structure specifying physical device
properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceProperties2` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDeviceProperties2 {
    VkStructureType               sType;
    void*                         pNext;
    VkPhysicalDeviceProperties    properties;
} VkPhysicalDeviceProperties2;
```

or the equivalent

``` c
// Provided by VK_KHR_get_physical_device_properties2
typedef VkPhysicalDeviceProperties2 VkPhysicalDeviceProperties2KHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `properties` is a
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)
  structure describing properties of the physical device. This structure
  is written with the same values as if it were written by
  [vkGetPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties.html).

## <a href="#_description" class="anchor"></a>Description

The `pNext` chain of this structure is used to extend the structure with
properties defined by extensions.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceProperties2-sType-sType"
  id="VUID-VkPhysicalDeviceProperties2-sType-sType"></a>
  VUID-VkPhysicalDeviceProperties2-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PROPERTIES_2`

- <a href="#VUID-VkPhysicalDeviceProperties2-pNext-pNext"
  id="VUID-VkPhysicalDeviceProperties2-pNext-pNext"></a>
  VUID-VkPhysicalDeviceProperties2-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html),
  [VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT.html),
  [VkPhysicalDeviceClusterCullingShaderPropertiesHUAWEI](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceClusterCullingShaderPropertiesHUAWEI.html),
  [VkPhysicalDeviceConservativeRasterizationPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceConservativeRasterizationPropertiesEXT.html),
  [VkPhysicalDeviceCooperativeMatrixPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCooperativeMatrixPropertiesKHR.html),
  [VkPhysicalDeviceCooperativeMatrixPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCooperativeMatrixPropertiesNV.html),
  [VkPhysicalDeviceCopyMemoryIndirectPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCopyMemoryIndirectPropertiesNV.html),
  [VkPhysicalDeviceCudaKernelLaunchPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCudaKernelLaunchPropertiesNV.html),
  [VkPhysicalDeviceCustomBorderColorPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCustomBorderColorPropertiesEXT.html),
  [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html),
  [VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferDensityMapPropertiesEXT.html),
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html),
  [VkPhysicalDeviceDescriptorIndexingProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorIndexingProperties.html),
  [VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDeviceGeneratedCommandsPropertiesNV.html),
  [VkPhysicalDeviceDiscardRectanglePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDiscardRectanglePropertiesEXT.html),
  [VkPhysicalDeviceDisplacementMicromapPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDisplacementMicromapPropertiesNV.html),
  [VkPhysicalDeviceDriverProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDriverProperties.html),
  [VkPhysicalDeviceDrmPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDrmPropertiesEXT.html),
  [VkPhysicalDeviceExtendedDynamicState3PropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExtendedDynamicState3PropertiesEXT.html),
  [VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExtendedSparseAddressSpacePropertiesNV.html),
  [VkPhysicalDeviceExternalFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalFormatResolvePropertiesANDROID.html),
  [VkPhysicalDeviceExternalMemoryHostPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalMemoryHostPropertiesEXT.html),
  [VkPhysicalDeviceFloatControlsProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFloatControlsProperties.html),
  [VkPhysicalDeviceFragmentDensityMap2PropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentDensityMap2PropertiesEXT.html),
  [VkPhysicalDeviceFragmentDensityMapOffsetPropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentDensityMapOffsetPropertiesQCOM.html),
  [VkPhysicalDeviceFragmentDensityMapPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentDensityMapPropertiesEXT.html),
  [VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentShaderBarycentricPropertiesKHR.html),
  [VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentShadingRateEnumsPropertiesNV.html),
  [VkPhysicalDeviceFragmentShadingRatePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentShadingRatePropertiesKHR.html),
  [VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceGraphicsPipelineLibraryPropertiesEXT.html),
  [VkPhysicalDeviceHostImageCopyPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceHostImageCopyPropertiesEXT.html),
  [VkPhysicalDeviceIDProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceIDProperties.html),
  [VkPhysicalDeviceImageAlignmentControlPropertiesMESA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageAlignmentControlPropertiesMESA.html),
  [VkPhysicalDeviceImageProcessing2PropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageProcessing2PropertiesQCOM.html),
  [VkPhysicalDeviceImageProcessingPropertiesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceImageProcessingPropertiesQCOM.html),
  [VkPhysicalDeviceInlineUniformBlockProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceInlineUniformBlockProperties.html),
  [VkPhysicalDeviceLayeredDriverPropertiesMSFT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredDriverPropertiesMSFT.html),
  [VkPhysicalDeviceLegacyVertexAttributesPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLegacyVertexAttributesPropertiesEXT.html),
  [VkPhysicalDeviceLineRasterizationPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLineRasterizationPropertiesKHR.html),
  [VkPhysicalDeviceMaintenance3Properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance3Properties.html),
  [VkPhysicalDeviceMaintenance4Properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance4Properties.html),
  [VkPhysicalDeviceMaintenance5PropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance5PropertiesKHR.html),
  [VkPhysicalDeviceMaintenance6PropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance6PropertiesKHR.html),
  [VkPhysicalDeviceMapMemoryPlacedPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMapMemoryPlacedPropertiesEXT.html),
  [VkPhysicalDeviceMemoryDecompressionPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryDecompressionPropertiesNV.html),
  [VkPhysicalDeviceMeshShaderPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesEXT.html),
  [VkPhysicalDeviceMeshShaderPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMeshShaderPropertiesNV.html),
  [VkPhysicalDeviceMultiDrawPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiDrawPropertiesEXT.html),
  [VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiviewPerViewAttributesPropertiesNVX.html),
  [VkPhysicalDeviceMultiviewProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMultiviewProperties.html),
  [VkPhysicalDeviceNestedCommandBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceNestedCommandBufferPropertiesEXT.html),
  [VkPhysicalDeviceOpacityMicromapPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceOpacityMicromapPropertiesEXT.html),
  [VkPhysicalDeviceOpticalFlowPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceOpticalFlowPropertiesNV.html),
  [VkPhysicalDevicePCIBusInfoPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePCIBusInfoPropertiesEXT.html),
  [VkPhysicalDevicePerformanceQueryPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePerformanceQueryPropertiesKHR.html),
  [VkPhysicalDevicePipelineRobustnessPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePipelineRobustnessPropertiesEXT.html),
  [VkPhysicalDevicePointClippingProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePointClippingProperties.html),
  [VkPhysicalDevicePortabilitySubsetPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetPropertiesKHR.html),
  [VkPhysicalDeviceProtectedMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProtectedMemoryProperties.html),
  [VkPhysicalDeviceProvokingVertexPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProvokingVertexPropertiesEXT.html),
  [VkPhysicalDevicePushDescriptorPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePushDescriptorPropertiesKHR.html),
  [VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingInvocationReorderPropertiesNV.html),
  [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html),
  [VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html),
  [VkPhysicalDeviceRenderPassStripedPropertiesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRenderPassStripedPropertiesARM.html),
  [VkPhysicalDeviceRobustness2PropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRobustness2PropertiesEXT.html),
  [VkPhysicalDeviceSampleLocationsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSampleLocationsPropertiesEXT.html),
  [VkPhysicalDeviceSamplerFilterMinmaxProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSamplerFilterMinmaxProperties.html),
  [VkPhysicalDeviceSchedulingControlsPropertiesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSchedulingControlsPropertiesARM.html),
  [VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderCoreBuiltinsPropertiesARM.html),
  [VkPhysicalDeviceShaderCoreProperties2AMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderCoreProperties2AMD.html),
  [VkPhysicalDeviceShaderCorePropertiesAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderCorePropertiesAMD.html),
  [VkPhysicalDeviceShaderCorePropertiesARM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderCorePropertiesARM.html),
  [VkPhysicalDeviceShaderEnqueuePropertiesAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderEnqueuePropertiesAMDX.html),
  [VkPhysicalDeviceShaderIntegerDotProductProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderIntegerDotProductProperties.html),
  [VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderModuleIdentifierPropertiesEXT.html),
  [VkPhysicalDeviceShaderObjectPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderObjectPropertiesEXT.html),
  [VkPhysicalDeviceShaderSMBuiltinsPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderSMBuiltinsPropertiesNV.html),
  [VkPhysicalDeviceShaderTileImagePropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderTileImagePropertiesEXT.html),
  [VkPhysicalDeviceShadingRateImagePropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShadingRateImagePropertiesNV.html),
  [VkPhysicalDeviceSubgroupProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSubgroupProperties.html),
  [VkPhysicalDeviceSubgroupSizeControlProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSubgroupSizeControlProperties.html),
  [VkPhysicalDeviceSubpassShadingPropertiesHUAWEI](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSubpassShadingPropertiesHUAWEI.html),
  [VkPhysicalDeviceTexelBufferAlignmentProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTexelBufferAlignmentProperties.html),
  [VkPhysicalDeviceTimelineSemaphoreProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTimelineSemaphoreProperties.html),
  [VkPhysicalDeviceTransformFeedbackPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceTransformFeedbackPropertiesEXT.html),
  [VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVertexAttributeDivisorPropertiesEXT.html),
  [VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVertexAttributeDivisorPropertiesKHR.html),
  [VkPhysicalDeviceVulkan11Properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan11Properties.html),
  [VkPhysicalDeviceVulkan12Properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan12Properties.html),
  or
  [VkPhysicalDeviceVulkan13Properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkan13Properties.html)

- <a href="#VUID-VkPhysicalDeviceProperties2-sType-unique"
  id="VUID-VkPhysicalDeviceProperties2-sType-unique"></a>
  VUID-VkPhysicalDeviceProperties2-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
[vkGetPhysicalDeviceProperties2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceProperties2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
