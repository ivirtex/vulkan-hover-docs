# VkPhysicalDeviceVulkan12Properties(3) Manual Page

## Name

VkPhysicalDeviceVulkan12Properties - Structure specifying physical device properties for functionality promoted to Vulkan 1.2



## [](#_c_specification)C Specification

The `VkPhysicalDeviceVulkan12Properties` structure is defined as:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceVulkan12Properties {
    VkStructureType                      sType;
    void*                                pNext;
    VkDriverId                           driverID;
    char                                 driverName[VK_MAX_DRIVER_NAME_SIZE];
    char                                 driverInfo[VK_MAX_DRIVER_INFO_SIZE];
    VkConformanceVersion                 conformanceVersion;
    VkShaderFloatControlsIndependence    denormBehaviorIndependence;
    VkShaderFloatControlsIndependence    roundingModeIndependence;
    VkBool32                             shaderSignedZeroInfNanPreserveFloat16;
    VkBool32                             shaderSignedZeroInfNanPreserveFloat32;
    VkBool32                             shaderSignedZeroInfNanPreserveFloat64;
    VkBool32                             shaderDenormPreserveFloat16;
    VkBool32                             shaderDenormPreserveFloat32;
    VkBool32                             shaderDenormPreserveFloat64;
    VkBool32                             shaderDenormFlushToZeroFloat16;
    VkBool32                             shaderDenormFlushToZeroFloat32;
    VkBool32                             shaderDenormFlushToZeroFloat64;
    VkBool32                             shaderRoundingModeRTEFloat16;
    VkBool32                             shaderRoundingModeRTEFloat32;
    VkBool32                             shaderRoundingModeRTEFloat64;
    VkBool32                             shaderRoundingModeRTZFloat16;
    VkBool32                             shaderRoundingModeRTZFloat32;
    VkBool32                             shaderRoundingModeRTZFloat64;
    uint32_t                             maxUpdateAfterBindDescriptorsInAllPools;
    VkBool32                             shaderUniformBufferArrayNonUniformIndexingNative;
    VkBool32                             shaderSampledImageArrayNonUniformIndexingNative;
    VkBool32                             shaderStorageBufferArrayNonUniformIndexingNative;
    VkBool32                             shaderStorageImageArrayNonUniformIndexingNative;
    VkBool32                             shaderInputAttachmentArrayNonUniformIndexingNative;
    VkBool32                             robustBufferAccessUpdateAfterBind;
    VkBool32                             quadDivergentImplicitLod;
    uint32_t                             maxPerStageDescriptorUpdateAfterBindSamplers;
    uint32_t                             maxPerStageDescriptorUpdateAfterBindUniformBuffers;
    uint32_t                             maxPerStageDescriptorUpdateAfterBindStorageBuffers;
    uint32_t                             maxPerStageDescriptorUpdateAfterBindSampledImages;
    uint32_t                             maxPerStageDescriptorUpdateAfterBindStorageImages;
    uint32_t                             maxPerStageDescriptorUpdateAfterBindInputAttachments;
    uint32_t                             maxPerStageUpdateAfterBindResources;
    uint32_t                             maxDescriptorSetUpdateAfterBindSamplers;
    uint32_t                             maxDescriptorSetUpdateAfterBindUniformBuffers;
    uint32_t                             maxDescriptorSetUpdateAfterBindUniformBuffersDynamic;
    uint32_t                             maxDescriptorSetUpdateAfterBindStorageBuffers;
    uint32_t                             maxDescriptorSetUpdateAfterBindStorageBuffersDynamic;
    uint32_t                             maxDescriptorSetUpdateAfterBindSampledImages;
    uint32_t                             maxDescriptorSetUpdateAfterBindStorageImages;
    uint32_t                             maxDescriptorSetUpdateAfterBindInputAttachments;
    VkResolveModeFlags                   supportedDepthResolveModes;
    VkResolveModeFlags                   supportedStencilResolveModes;
    VkBool32                             independentResolveNone;
    VkBool32                             independentResolve;
    VkBool32                             filterMinmaxSingleComponentFormats;
    VkBool32                             filterMinmaxImageComponentMapping;
    uint64_t                             maxTimelineSemaphoreValueDifference;
    VkSampleCountFlags                   framebufferIntegerColorSampleCounts;
} VkPhysicalDeviceVulkan12Properties;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- `driverID` is a unique identifier for the driver of the physical device.
- `driverName` is an array of `VK_MAX_DRIVER_NAME_SIZE` `char` containing a null-terminated UTF-8 string which is the name of the driver.
- `driverInfo` is an array of `VK_MAX_DRIVER_INFO_SIZE` `char` containing a null-terminated UTF-8 string with additional information about the driver.
- `conformanceVersion` is the latest version of the Vulkan conformance test that the implementor has successfully tested this driver against prior to release (see [VkConformanceVersion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConformanceVersion.html)).
- []()`denormBehaviorIndependence` is a [VkShaderFloatControlsIndependence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderFloatControlsIndependence.html) value indicating whether, and how, denorm behavior can be set independently for different bit widths.
- []()`roundingModeIndependence` is a [VkShaderFloatControlsIndependence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderFloatControlsIndependence.html) value indicating whether, and how, rounding modes can be set independently for different bit widths.
- []()`shaderSignedZeroInfNanPreserveFloat16` is a boolean value indicating whether sign of a zero, Nans and ±∞ **can** be preserved in 16-bit floating-point computations. It also indicates whether the `SignedZeroInfNanPreserve` execution mode **can** be used for 16-bit floating-point types.
- []()`shaderSignedZeroInfNanPreserveFloat32` is a boolean value indicating whether sign of a zero, Nans and ±∞ **can** be preserved in 32-bit floating-point computations. It also indicates whether the `SignedZeroInfNanPreserve` execution mode **can** be used for 32-bit floating-point types.
- []()`shaderSignedZeroInfNanPreserveFloat64` is a boolean value indicating whether sign of a zero, Nans and ±∞ **can** be preserved in 64-bit floating-point computations. It also indicates whether the `SignedZeroInfNanPreserve` execution mode **can** be used for 64-bit floating-point types.
- []()`shaderDenormPreserveFloat16` is a boolean value indicating whether denormals **can** be preserved in 16-bit floating-point computations. It also indicates whether the `DenormPreserve` execution mode **can** be used for 16-bit floating-point types.
- []()`shaderDenormPreserveFloat32` is a boolean value indicating whether denormals **can** be preserved in 32-bit floating-point computations. It also indicates whether the `DenormPreserve` execution mode **can** be used for 32-bit floating-point types.
- []()`shaderDenormPreserveFloat64` is a boolean value indicating whether denormals **can** be preserved in 64-bit floating-point computations. It also indicates whether the `DenormPreserve` execution mode **can** be used for 64-bit floating-point types.
- []()`shaderDenormFlushToZeroFloat16` is a boolean value indicating whether denormals **can** be flushed to zero in 16-bit floating-point computations. It also indicates whether the `DenormFlushToZero` execution mode **can** be used for 16-bit floating-point types.
- []()`shaderDenormFlushToZeroFloat32` is a boolean value indicating whether denormals **can** be flushed to zero in 32-bit floating-point computations. It also indicates whether the `DenormFlushToZero` execution mode **can** be used for 32-bit floating-point types.
- []()`shaderDenormFlushToZeroFloat64` is a boolean value indicating whether denormals **can** be flushed to zero in 64-bit floating-point computations. It also indicates whether the `DenormFlushToZero` execution mode **can** be used for 64-bit floating-point types.
- []()`shaderRoundingModeRTEFloat16` is a boolean value indicating whether an implementation supports the round-to-nearest-even rounding mode for 16-bit floating-point arithmetic and conversion instructions. It also indicates whether the `RoundingModeRTE` execution mode **can** be used for 16-bit floating-point types.
- []()`shaderRoundingModeRTEFloat32` is a boolean value indicating whether an implementation supports the round-to-nearest-even rounding mode for 32-bit floating-point arithmetic and conversion instructions. It also indicates whether the `RoundingModeRTE` execution mode **can** be used for 32-bit floating-point types.
- []()`shaderRoundingModeRTEFloat64` is a boolean value indicating whether an implementation supports the round-to-nearest-even rounding mode for 64-bit floating-point arithmetic and conversion instructions. It also indicates whether the `RoundingModeRTE` execution mode **can** be used for 64-bit floating-point types.
- []()`shaderRoundingModeRTZFloat16` is a boolean value indicating whether an implementation supports the round-towards-zero rounding mode for 16-bit floating-point arithmetic and conversion instructions. It also indicates whether the `RoundingModeRTZ` execution mode **can** be used for 16-bit floating-point types.
- []()`shaderRoundingModeRTZFloat32` is a boolean value indicating whether an implementation supports the round-towards-zero rounding mode for 32-bit floating-point arithmetic and conversion instructions. It also indicates whether the `RoundingModeRTZ` execution mode **can** be used for 32-bit floating-point types.
- []()`shaderRoundingModeRTZFloat64` is a boolean value indicating whether an implementation supports the round-towards-zero rounding mode for 64-bit floating-point arithmetic and conversion instructions. It also indicates whether the `RoundingModeRTZ` execution mode **can** be used for 64-bit floating-point types.
- []()`maxUpdateAfterBindDescriptorsInAllPools` is the maximum number of descriptors (summed over all descriptor types) that **can** be created across all pools that are created with the `VK_DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT` bit set. Pool creation **may** fail when this limit is exceeded, or when the space this limit represents is unable to satisfy a pool creation due to fragmentation.
- []()`shaderUniformBufferArrayNonUniformIndexingNative` is a boolean value indicating whether uniform buffer descriptors natively support non-uniform indexing. If this is `VK_FALSE`, then a single dynamic instance of an instruction that non-uniformly indexes an array of uniform buffers **may** execute multiple times in order to access all the descriptors.
- []()`shaderSampledImageArrayNonUniformIndexingNative` is a boolean value indicating whether sampler and image descriptors natively support non-uniform indexing. If this is `VK_FALSE`, then a single dynamic instance of an instruction that non-uniformly indexes an array of samplers or images **may** execute multiple times in order to access all the descriptors.
- []()`shaderStorageBufferArrayNonUniformIndexingNative` is a boolean value indicating whether storage buffer descriptors natively support non-uniform indexing. If this is `VK_FALSE`, then a single dynamic instance of an instruction that non-uniformly indexes an array of storage buffers **may** execute multiple times in order to access all the descriptors.
- []()`shaderStorageImageArrayNonUniformIndexingNative` is a boolean value indicating whether storage image descriptors natively support non-uniform indexing. If this is `VK_FALSE`, then a single dynamic instance of an instruction that non-uniformly indexes an array of storage images **may** execute multiple times in order to access all the descriptors.
- []()`shaderInputAttachmentArrayNonUniformIndexingNative` is a boolean value indicating whether input attachment descriptors natively support non-uniform indexing. If this is `VK_FALSE`, then a single dynamic instance of an instruction that non-uniformly indexes an array of input attachments **may** execute multiple times in order to access all the descriptors.
- []()`robustBufferAccessUpdateAfterBind` is a boolean value indicating whether [`robustBufferAccess`](#features-robustBufferAccess) **can** be enabled on a device simultaneously with `descriptorBindingUniformBufferUpdateAfterBind`, `descriptorBindingStorageBufferUpdateAfterBind`, `descriptorBindingUniformTexelBufferUpdateAfterBind`, and/or `descriptorBindingStorageTexelBufferUpdateAfterBind`. If this is `VK_FALSE`, then either `robustBufferAccess` **must** be disabled or all of these update-after-bind features **must** be disabled. Similarly, if this property is `VK_FALSE`, robustness **must** not be enabled through the [VkPipelineRobustnessCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessCreateInfo.html) mechanism.
- []()`quadDivergentImplicitLod` is a boolean value indicating whether implicit LOD calculations for image operations have well-defined results when the image and/or sampler objects used for the instruction are not uniform within a quad. See [Derivative Image Operations](#textures-derivative-image-operations).
- []()`maxPerStageDescriptorUpdateAfterBindSamplers` is similar to `maxPerStageDescriptorSamplers` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxPerStageDescriptorUpdateAfterBindUniformBuffers` is similar to `maxPerStageDescriptorUniformBuffers` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxPerStageDescriptorUpdateAfterBindStorageBuffers` is similar to `maxPerStageDescriptorStorageBuffers` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxPerStageDescriptorUpdateAfterBindSampledImages` is similar to `maxPerStageDescriptorSampledImages` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxPerStageDescriptorUpdateAfterBindStorageImages` is similar to `maxPerStageDescriptorStorageImages` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxPerStageDescriptorUpdateAfterBindInputAttachments` is similar to `maxPerStageDescriptorInputAttachments` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxPerStageUpdateAfterBindResources` is similar to `maxPerStageResources` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxDescriptorSetUpdateAfterBindSamplers` is similar to `maxDescriptorSetSamplers` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxDescriptorSetUpdateAfterBindUniformBuffers` is similar to `maxDescriptorSetUniformBuffers` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxDescriptorSetUpdateAfterBindUniformBuffersDynamic` is similar to `maxDescriptorSetUniformBuffersDynamic` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set. While an application **can** allocate dynamic uniform buffer descriptors from a pool created with the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT`, bindings for these descriptors **must** not be present in any descriptor set layout that includes bindings created with `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`.
- []()`maxDescriptorSetUpdateAfterBindStorageBuffers` is similar to `maxDescriptorSetStorageBuffers` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxDescriptorSetUpdateAfterBindStorageBuffersDynamic` is similar to `maxDescriptorSetStorageBuffersDynamic` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set. While an application **can** allocate dynamic storage buffer descriptors from a pool created with the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT`, bindings for these descriptors **must** not be present in any descriptor set layout that includes bindings created with `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT`.
- []()`maxDescriptorSetUpdateAfterBindSampledImages` is similar to `maxDescriptorSetSampledImages` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxDescriptorSetUpdateAfterBindStorageImages` is similar to `maxDescriptorSetStorageImages` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxDescriptorSetUpdateAfterBindInputAttachments` is similar to `maxDescriptorSetInputAttachments` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`supportedDepthResolveModes` is a bitmask of [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlagBits.html) indicating the set of supported depth resolve modes. `VK_RESOLVE_MODE_SAMPLE_ZERO_BIT` **must** be included in the set but implementations **may** support additional modes.
- []()`supportedStencilResolveModes` is a bitmask of [VkResolveModeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlagBits.html) indicating the set of supported stencil resolve modes. `VK_RESOLVE_MODE_SAMPLE_ZERO_BIT` **must** be included in the set but implementations **may** support additional modes. `VK_RESOLVE_MODE_AVERAGE_BIT` **must** not be included in the set.
- []()`independentResolveNone` is `VK_TRUE` if the implementation supports setting the depth and stencil resolve modes to different values when one of those modes is `VK_RESOLVE_MODE_NONE`. Otherwise the implementation only supports setting both modes to the same value.
- []()`independentResolve` is `VK_TRUE` if the implementation supports all combinations of the supported depth and stencil resolve modes, including setting either depth or stencil resolve mode to `VK_RESOLVE_MODE_NONE`. An implementation that supports `independentResolve` **must** also support `independentResolveNone`.
- []()`filterMinmaxSingleComponentFormats` is a boolean value indicating whether a minimum set of required formats support min/max filtering.
- []()`filterMinmaxImageComponentMapping` is a boolean value indicating whether the implementation supports non-identity component mapping of the image when doing min/max filtering.
- []()`maxTimelineSemaphoreValueDifference` indicates the maximum difference allowed by the implementation between the current value of a timeline semaphore and any pending signal or wait operations.
- []()`framebufferIntegerColorSampleCounts` is a bitmask of [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) indicating the color sample counts that are supported for all framebuffer color attachments with integer formats.

If the `VkPhysicalDeviceVulkan12Properties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

These properties correspond to Vulkan 1.2 functionality.

The members of `VkPhysicalDeviceVulkan12Properties` **must** have the same values as the corresponding members of [VkPhysicalDeviceDriverProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDriverProperties.html), [VkPhysicalDeviceFloatControlsProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFloatControlsProperties.html), [VkPhysicalDeviceDescriptorIndexingProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDescriptorIndexingProperties.html), [VkPhysicalDeviceDepthStencilResolveProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceDepthStencilResolveProperties.html), [VkPhysicalDeviceSamplerFilterMinmaxProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSamplerFilterMinmaxProperties.html), and [VkPhysicalDeviceTimelineSemaphoreProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceTimelineSemaphoreProperties.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceVulkan12Properties-sType-sType)VUID-VkPhysicalDeviceVulkan12Properties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_PROPERTIES`

## [](#_see_also)See Also

[VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkConformanceVersion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConformanceVersion.html), [VkDriverId](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDriverId.html), [VkResolveModeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlags.html), [VkSampleCountFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlags.html), [VkShaderFloatControlsIndependence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderFloatControlsIndependence.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceVulkan12Properties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0