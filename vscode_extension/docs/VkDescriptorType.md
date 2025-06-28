# VkDescriptorType(3) Manual Page

## Name

VkDescriptorType - Specifies the type of a descriptor in a descriptor set



## [](#_c_specification)C Specification

The type of descriptors in a descriptor set is specified by [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html)::`descriptorType`, which **must** be one of the values:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkDescriptorType {
    VK_DESCRIPTOR_TYPE_SAMPLER = 0,
    VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER = 1,
    VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE = 2,
    VK_DESCRIPTOR_TYPE_STORAGE_IMAGE = 3,
    VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER = 4,
    VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER = 5,
    VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER = 6,
    VK_DESCRIPTOR_TYPE_STORAGE_BUFFER = 7,
    VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC = 8,
    VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC = 9,
    VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT = 10,
  // Provided by VK_VERSION_1_3
    VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK = 1000138000,
  // Provided by VK_KHR_acceleration_structure
    VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR = 1000150000,
  // Provided by VK_NV_ray_tracing
    VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV = 1000165000,
  // Provided by VK_QCOM_image_processing
    VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM = 1000440000,
  // Provided by VK_QCOM_image_processing
    VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM = 1000440001,
  // Provided by VK_ARM_tensors
    VK_DESCRIPTOR_TYPE_TENSOR_ARM = 1000460000,
  // Provided by VK_EXT_mutable_descriptor_type
    VK_DESCRIPTOR_TYPE_MUTABLE_EXT = 1000351000,
  // Provided by VK_NV_partitioned_acceleration_structure
    VK_DESCRIPTOR_TYPE_PARTITIONED_ACCELERATION_STRUCTURE_NV = 1000570000,
  // Provided by VK_EXT_inline_uniform_block
    VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK_EXT = VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK,
  // Provided by VK_VALVE_mutable_descriptor_type
    VK_DESCRIPTOR_TYPE_MUTABLE_VALVE = VK_DESCRIPTOR_TYPE_MUTABLE_EXT,
} VkDescriptorType;
```

## [](#_description)Description

- `VK_DESCRIPTOR_TYPE_SAMPLER` specifies a [sampler descriptor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-sampler).
- `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` specifies a [combined image sampler descriptor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-combinedimagesampler).
- `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` specifies a [sampled image descriptor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-sampledimage).
- `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` specifies a [storage image descriptor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storageimage).
- `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` specifies a [uniform texel buffer descriptor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-uniformtexelbuffer).
- `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` specifies a [storage texel buffer descriptor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagetexelbuffer).
- `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` specifies a [uniform buffer descriptor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-uniformbuffer).
- `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` specifies a [storage buffer descriptor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagebuffer).
- `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` specifies a [dynamic uniform buffer descriptor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-uniformbufferdynamic).
- `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` specifies a [dynamic storage buffer descriptor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagebufferdynamic).
- `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` specifies an [input attachment descriptor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-inputattachment).
- `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` specifies an [inline uniform block](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-inlineuniformblock).
- `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` specifies a [descriptor of mutable type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-mutable).
- `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM` specifies a [sampled weight image descriptor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-weightimage).
- `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM` specifies a [block matching image descriptor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-blockmatch).
- `VK_DESCRIPTOR_TYPE_TENSOR_ARM` specifies a [storage tensor descriptor](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storagetensor).

When a descriptor set is updated via elements of [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html), members of `pImageInfo`, `pBufferInfo` and `pTexelBufferView` are only accessed by the implementation when they correspond to descriptor type being defined - otherwise they are ignored. The members accessed are as follows for each descriptor type:

- For `VK_DESCRIPTOR_TYPE_SAMPLER`, only the `sampler` member of each element of [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html)::`pImageInfo` is accessed.
- For `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, only the `imageView` and `imageLayout` members of each element of [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html)::`pImageInfo` are accessed.
- For `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, all members of each element of [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html)::`pImageInfo` are accessed.
- For `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`, `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`, `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`, or `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`, all members of each element of [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html)::`pBufferInfo` are accessed.
- For `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` or `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`, each element of [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html)::`pTexelBufferView` is accessed.

When updating descriptors with a `descriptorType` of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, none of the `pImageInfo`, `pBufferInfo`, or `pTexelBufferView` members are accessed, instead the source data of the descriptor update operation is taken from the [VkWriteDescriptorSetInlineUniformBlock](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetInlineUniformBlock.html) structure in the `pNext` chain of `VkWriteDescriptorSet`. When updating descriptors with a `descriptorType` of `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR`, none of the `pImageInfo`, `pBufferInfo`, or `pTexelBufferView` members are accessed, instead the source data of the descriptor update operation is taken from the [VkWriteDescriptorSetAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetAccelerationStructureKHR.html) structure in the `pNext` chain of `VkWriteDescriptorSet`. When updating descriptors with a `descriptorType` of `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV`, none of the `pImageInfo`, `pBufferInfo`, or `pTexelBufferView` members are accessed, instead the source data of the descriptor update operation is taken from the [VkWriteDescriptorSetAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetAccelerationStructureNV.html) structure in the `pNext` chain of `VkWriteDescriptorSet`. When updating descriptors with a `descriptorType` of `VK_DESCRIPTOR_TYPE_TENSOR_ARM`, none of the `pImageInfo`, `pBufferInfo`, or `pTexelBufferView` members are accessed, instead the source data of the descriptor update operation is taken from the instance of [VkWriteDescriptorSetTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetTensorARM.html) in the `pNext` chain of [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html).

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorGetInfoEXT.html), [VkDescriptorPoolSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolSize.html), [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBinding.html), [VkDescriptorUpdateTemplateEntry](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorUpdateTemplateEntry.html), [VkImageViewHandleInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewHandleInfoNVX.html), [VkMutableDescriptorTypeListEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMutableDescriptorTypeListEXT.html), [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSet.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorType)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0