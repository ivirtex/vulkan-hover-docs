# VkPhysicalDeviceTensorPropertiesARM(3) Manual Page

## Name

VkPhysicalDeviceTensorPropertiesARM - Structure describing the tensor properties of a physical device



## [](#_c_specification)C Specification

The `VkPhysicalDeviceTensorPropertiesARM` structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkPhysicalDeviceTensorPropertiesARM {
    VkStructureType       sType;
    void*                 pNext;
    uint32_t              maxTensorDimensionCount;
    uint64_t              maxTensorElements;
    uint64_t              maxPerDimensionTensorElements;
    int64_t               maxTensorStride;
    uint64_t              maxTensorSize;
    uint32_t              maxTensorShaderAccessArrayLength;
    uint32_t              maxTensorShaderAccessSize;
    uint32_t              maxDescriptorSetStorageTensors;
    uint32_t              maxPerStageDescriptorSetStorageTensors;
    uint32_t              maxDescriptorSetUpdateAfterBindStorageTensors;
    uint32_t              maxPerStageDescriptorUpdateAfterBindStorageTensors;
    VkBool32              shaderStorageTensorArrayNonUniformIndexingNative;
    VkShaderStageFlags    shaderTensorSupportedStages;
} VkPhysicalDeviceTensorPropertiesARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`maxTensorDimensionCount` is the maximum number of dimensions that can be specified in the `dimensionCount` member of [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html).
- []()`maxTensorElements` is the maximum number of data elements in a created tensor as specified in the [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html) of [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html). The number of data elements in a tensor is computed as the product of `pDimensions`\[i] for all 0 ≤ i ≤ dimensionCount-1.
- []()`maxPerDimensionTensorElements` is the maximum number of data elements alongside any dimension of a tensor.
- []()`maxTensorStride` is the maximum value for a tensor stride that can be used in [VkTensorDescriptionARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorDescriptionARM.html)::`pStrides`.
- []()`maxTensorSize` is the maximum size, in bytes, of a tensor.
- []()`maxTensorShaderAccessArrayLength` is the maximum number of elements in an array returned by `OpTensoReadARM` or consumed by `OpTensorWriteARM`.
- []()`maxTensorShaderAccessSize` is the maximum size in bytes of the data that can be read from a tensor with `OpTensorReadARM` or written to a tensor with `OpTensorWriteARM`.
- []()`maxDescriptorSetStorageTensors` is the maximum number of tensors that **can** be included in descriptor bindings in a pipeline layout across all pipeline shader stages and descriptor set numbers. Descriptors with a type of `VK_DESCRIPTOR_TYPE_TENSOR_ARM` count against this limit.
- []()`maxPerStageDescriptorSetStorageTensors` is the maximum number of tensors that **can** be accessible to a single shader stage in a pipeline layout. Descriptors with a type of `VK_DESCRIPTOR_TYPE_TENSOR_ARM` count against this limit. A descriptor is accessible to a pipeline shader stage when the `stageFlags` member of the [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBinding.html) structure has the bit for that shader stage set.
- []()`maxDescriptorSetUpdateAfterBindStorageTensors` is similar to `maxDescriptorSetStorageTensors` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`maxPerStageDescriptorUpdateAfterBindStorageTensors` is similar to `maxPerStageDescriptorSetStorageTensors` but counts descriptors from descriptor sets created with or without the `VK_DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT` bit set.
- []()`shaderStorageTensorArrayNonUniformIndexingNative` is a boolean value indicating whether storage tensor descriptors natively support nonuniform indexing. If this is `VK_FALSE`, then a single dynamic instance of an instruction that nonuniformly indexes an array of storage buffers may execute multiple times in order to access all the descriptors.
- []()`shaderTensorSupportedStages` is a bitfield of [VkShaderStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlagBits.html) describing the shader stages that **can** access tensor resources. `shaderTensorSupportedStages` will have the `VK_SHADER_STAGE_COMPUTE_BIT` bit set if any of the physical device’s queues support `VK_QUEUE_COMPUTE_BIT`.

## [](#_description)Description

If the `VkPhysicalDeviceTensorPropertiesARM` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceTensorPropertiesARM-sType-sType)VUID-VkPhysicalDeviceTensorPropertiesARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TENSOR_PROPERTIES_ARM`

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceTensorPropertiesARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0