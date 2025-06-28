# VkPhysicalDeviceVulkan13Properties(3) Manual Page

## Name

VkPhysicalDeviceVulkan13Properties - Structure specifying physical device properties for functionality promoted to Vulkan 1.3



## [](#_c_specification)C Specification

The `VkPhysicalDeviceVulkan13Properties` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkPhysicalDeviceVulkan13Properties {
    VkStructureType       sType;
    void*                 pNext;
    uint32_t              minSubgroupSize;
    uint32_t              maxSubgroupSize;
    uint32_t              maxComputeWorkgroupSubgroups;
    VkShaderStageFlags    requiredSubgroupSizeStages;
    uint32_t              maxInlineUniformBlockSize;
    uint32_t              maxPerStageDescriptorInlineUniformBlocks;
    uint32_t              maxPerStageDescriptorUpdateAfterBindInlineUniformBlocks;
    uint32_t              maxDescriptorSetInlineUniformBlocks;
    uint32_t              maxDescriptorSetUpdateAfterBindInlineUniformBlocks;
    uint32_t              maxInlineUniformTotalSize;
    VkBool32              integerDotProduct8BitUnsignedAccelerated;
    VkBool32              integerDotProduct8BitSignedAccelerated;
    VkBool32              integerDotProduct8BitMixedSignednessAccelerated;
    VkBool32              integerDotProduct4x8BitPackedUnsignedAccelerated;
    VkBool32              integerDotProduct4x8BitPackedSignedAccelerated;
    VkBool32              integerDotProduct4x8BitPackedMixedSignednessAccelerated;
    VkBool32              integerDotProduct16BitUnsignedAccelerated;
    VkBool32              integerDotProduct16BitSignedAccelerated;
    VkBool32              integerDotProduct16BitMixedSignednessAccelerated;
    VkBool32              integerDotProduct32BitUnsignedAccelerated;
    VkBool32              integerDotProduct32BitSignedAccelerated;
    VkBool32              integerDotProduct32BitMixedSignednessAccelerated;
    VkBool32              integerDotProduct64BitUnsignedAccelerated;
    VkBool32              integerDotProduct64BitSignedAccelerated;
    VkBool32              integerDotProduct64BitMixedSignednessAccelerated;
    VkBool32              integerDotProductAccumulatingSaturating8BitUnsignedAccelerated;
    VkBool32              integerDotProductAccumulatingSaturating8BitSignedAccelerated;
    VkBool32              integerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated;
    VkBool32              integerDotProductAccumulatingSaturating4x8BitPackedUnsignedAccelerated;
    VkBool32              integerDotProductAccumulatingSaturating4x8BitPackedSignedAccelerated;
    VkBool32              integerDotProductAccumulatingSaturating4x8BitPackedMixedSignednessAccelerated;
    VkBool32              integerDotProductAccumulatingSaturating16BitUnsignedAccelerated;
    VkBool32              integerDotProductAccumulatingSaturating16BitSignedAccelerated;
    VkBool32              integerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated;
    VkBool32              integerDotProductAccumulatingSaturating32BitUnsignedAccelerated;
    VkBool32              integerDotProductAccumulatingSaturating32BitSignedAccelerated;
    VkBool32              integerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated;
    VkBool32              integerDotProductAccumulatingSaturating64BitUnsignedAccelerated;
    VkBool32              integerDotProductAccumulatingSaturating64BitSignedAccelerated;
    VkBool32              integerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated;
    VkDeviceSize          storageTexelBufferOffsetAlignmentBytes;
    VkBool32              storageTexelBufferOffsetSingleTexelAlignment;
    VkDeviceSize          uniformTexelBufferOffsetAlignmentBytes;
    VkBool32              uniformTexelBufferOffsetSingleTexelAlignment;
    VkDeviceSize          maxBufferSize;
} VkPhysicalDeviceVulkan13Properties;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`minSubgroupSize` is the minimum subgroup size supported by this device. `minSubgroupSize` is at least one if any of the physical device’s queues support `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT`. `minSubgroupSize` is a power-of-two. `minSubgroupSize` is less than or equal to `maxSubgroupSize`. `minSubgroupSize` is less than or equal to [`subgroupSize`](#limits-subgroupSize).
- []()`maxSubgroupSize` is the maximum subgroup size supported by this device. `maxSubgroupSize` is at least one if any of the physical device’s queues support `VK_QUEUE_GRAPHICS_BIT` or `VK_QUEUE_COMPUTE_BIT`. `maxSubgroupSize` is a power-of-two. `maxSubgroupSize` is greater than or equal to `minSubgroupSize`. `maxSubgroupSize` is greater than or equal to [`subgroupSize`](#limits-subgroupSize).
- []()`maxComputeWorkgroupSubgroups` is the maximum number of subgroups supported by the implementation within a workgroup.
- []()`requiredSubgroupSizeStages` is a bitfield of what shader stages support having a required subgroup size specified.
- []()`maxInlineUniformBlockSize` is the maximum size in bytes of an [inline uniform block](#descriptorsets-inlineuniformblock) binding.
- []()`maxPerStageDescriptorInlineUniformBlocks` is the maximum number of inline uniform block bindings that **can** be accessible to a single shader stage in a pipeline layout. Descriptor bindings with a descriptor type of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` count against this limit. Only descriptor bindings in descriptor set layouts created without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set count against this limit.
- []()`maxPerStageDescriptorUpdateAfterBindInlineUniformBlocks` is similar to `maxPerStageDescriptorInlineUniformBlocks` but counts descriptor bindings from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxDescriptorSetInlineUniformBlocks` is the maximum number of inline uniform block bindings that **can** be included in descriptor bindings in a pipeline layout across all pipeline shader stages and descriptor set numbers. Descriptor bindings with a descriptor type of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` count against this limit. Only descriptor bindings in descriptor set layouts created without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set count against this limit.
- []()`maxDescriptorSetUpdateAfterBindInlineUniformBlocks` is similar to `maxDescriptorSetInlineUniformBlocks` but counts descriptor bindings from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxInlineUniformTotalSize` is the maximum total size in bytes of all inline uniform block bindings, across all pipeline shader stages and descriptor set numbers, that **can** be included in a pipeline layout. Descriptor bindings with a descriptor type of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` count against this limit.
- `integerDotProduct8BitUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit unsigned dot product operations using the `OpUDotKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct8BitSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit signed dot product operations using the `OpSDotKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct8BitMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit mixed signedness dot product operations using the `OpSUDotKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct4x8BitPackedUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit unsigned dot product operations from operands packed into 32-bit integers using the `OpUDotKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct4x8BitPackedSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit signed dot product operations from operands packed into 32-bit integers using the `OpSDotKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct4x8BitPackedMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit mixed signedness dot product operations from operands packed into 32-bit integers using the `OpSUDotKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct16BitUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 16-bit unsigned dot product operations using the `OpUDotKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct16BitSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 16-bit signed dot product operations using the `OpSDotKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct16BitMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 16-bit mixed signedness dot product operations using the `OpSUDotKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct32BitUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 32-bit unsigned dot product operations using the `OpUDotKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct32BitSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 32-bit signed dot product operations using the `OpSDotKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct32BitMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 32-bit mixed signedness dot product operations using the `OpSUDotKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct64BitUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 64-bit unsigned dot product operations using the `OpUDotKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct64BitSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 64-bit signed dot product operations using the `OpSDotKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProduct64BitMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 64-bit mixed signedness dot product operations using the `OpSUDotKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating8BitUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit unsigned accumulating saturating dot product operations using the `OpUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating8BitSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit signed accumulating saturating dot product operations using the `OpSDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit mixed signedness accumulating saturating dot product operations using the `OpSUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating4x8BitPackedUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit unsigned accumulating saturating dot product operations from operands packed into 32-bit integers using the `OpUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating4x8BitPackedSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit signed accumulating saturating dot product operations from operands packed into 32-bit integers using the `OpSDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating4x8BitPackedMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 8-bit mixed signedness accumulating saturating dot product operations from operands packed into 32-bit integers using the `OpSUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating16BitUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 16-bit unsigned accumulating saturating dot product operations using the `OpUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating16BitSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 16-bit signed accumulating saturating dot product operations using the `OpSDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 16-bit mixed signedness accumulating saturating dot product operations using the `OpSUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating32BitUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 32-bit unsigned accumulating saturating dot product operations using the `OpUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating32BitSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 32-bit signed accumulating saturating dot product operations using the `OpSDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 32-bit mixed signedness accumulating saturating dot product operations using the `OpSUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating64BitUnsignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 64-bit unsigned accumulating saturating dot product operations using the `OpUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating64BitSignedAccelerated` is a boolean that will be `VK_TRUE` if the support for 64-bit signed accumulating saturating dot product operations using the `OpSDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- `integerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated` is a boolean that will be `VK_TRUE` if the support for 64-bit mixed signedness accumulating saturating dot product operations using the `OpSUDotAccSatKHR` SPIR-V instruction is accelerated [as defined below](#devsandqueues-integer-dot-product-accelerated).
- []()`storageTexelBufferOffsetAlignmentBytes` is a byte alignment that is sufficient for a storage texel buffer of any format. The value **must** be a power of two.
- []()`storageTexelBufferOffsetSingleTexelAlignment` indicates whether single texel alignment is sufficient for a storage texel buffer of any format.
- []()`uniformTexelBufferOffsetAlignmentBytes` is a byte alignment that is sufficient for a uniform texel buffer of any format. The value **must** be a power of two.
- []()`uniformTexelBufferOffsetSingleTexelAlignment` indicates whether single texel alignment is sufficient for a uniform texel buffer of any format.
- []()`maxBufferSize` is the maximum size `VkBuffer` that **can** be created.

If the `VkPhysicalDeviceVulkan13Properties` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

These properties correspond to Vulkan 1.3 functionality.

The members of `VkPhysicalDeviceVulkan13Properties` **must** have the same values as the corresponding members of [VkPhysicalDeviceInlineUniformBlockProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceInlineUniformBlockProperties.html) and [VkPhysicalDeviceSubgroupSizeControlProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupSizeControlProperties.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceVulkan13Properties-sType-sType)VUID-VkPhysicalDeviceVulkan13Properties-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_3_PROPERTIES`

## [](#_see_also)See Also

[VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceVulkan13Properties)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0