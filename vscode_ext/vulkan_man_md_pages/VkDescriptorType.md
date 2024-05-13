# VkDescriptorType(3) Manual Page

## Name

VkDescriptorType - Specifies the type of a descriptor in a descriptor
set



## <a href="#_c_specification" class="anchor"></a>C Specification

The type of descriptors in a descriptor set is specified by
[VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html)::`descriptorType`,
which **must** be one of the values:

``` c
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
  // Provided by VK_EXT_mutable_descriptor_type
    VK_DESCRIPTOR_TYPE_MUTABLE_EXT = 1000351000,
  // Provided by VK_EXT_inline_uniform_block
    VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK_EXT = VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK,
  // Provided by VK_VALVE_mutable_descriptor_type
    VK_DESCRIPTOR_TYPE_MUTABLE_VALVE = VK_DESCRIPTOR_TYPE_MUTABLE_EXT,
} VkDescriptorType;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DESCRIPTOR_TYPE_SAMPLER` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-sampler"
  target="_blank" rel="noopener">sampler descriptor</a>.

- `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-combinedimagesampler"
  target="_blank" rel="noopener">combined image sampler descriptor</a>.

- `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-sampledimage"
  target="_blank" rel="noopener">sampled image descriptor</a>.

- `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storageimage"
  target="_blank" rel="noopener">storage image descriptor</a>.

- `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformtexelbuffer"
  target="_blank" rel="noopener">uniform texel buffer descriptor</a>.

- `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagetexelbuffer"
  target="_blank" rel="noopener">storage texel buffer descriptor</a>.

- `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformbuffer"
  target="_blank" rel="noopener">uniform buffer descriptor</a>.

- `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagebuffer"
  target="_blank" rel="noopener">storage buffer descriptor</a>.

- `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-uniformbufferdynamic"
  target="_blank" rel="noopener">dynamic uniform buffer descriptor</a>.

- `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storagebufferdynamic"
  target="_blank" rel="noopener">dynamic storage buffer descriptor</a>.

- `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` specifies an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-inputattachment"
  target="_blank" rel="noopener">input attachment descriptor</a>.

- `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` specifies an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-inlineuniformblock"
  target="_blank" rel="noopener">inline uniform block</a>.

- `VK_DESCRIPTOR_TYPE_MUTABLE_EXT` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-mutable"
  target="_blank" rel="noopener">descriptor of mutable type</a>.

- `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-weightimage"
  target="_blank" rel="noopener">sampled weight image descriptor</a>.

- `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM` specifies a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-blockmatch"
  target="_blank" rel="noopener">block matching image descriptor</a>.

When a descriptor set is updated via elements of
[VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html), members of
`pImageInfo`, `pBufferInfo` and `pTexelBufferView` are only accessed by
the implementation when they correspond to descriptor type being
defined - otherwise they are ignored. The members accessed are as
follows for each descriptor type:

- For `VK_DESCRIPTOR_TYPE_SAMPLER`, only the `sampler` member of each
  element of
  [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html)::`pImageInfo` is
  accessed.

- For `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`,
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, only the `imageView` and
  `imageLayout` members of each element of
  [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html)::`pImageInfo` are
  accessed.

- For `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, all members of each
  element of
  [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html)::`pImageInfo` are
  accessed.

- For `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`,
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`,
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`, or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`, all members of each
  element of
  [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html)::`pBufferInfo` are
  accessed.

- For `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` or
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`, each element of
  [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html)::`pTexelBufferView`
  is accessed.

When updating descriptors with a `descriptorType` of
`VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, none of the `pImageInfo`,
`pBufferInfo`, or `pTexelBufferView` members are accessed, instead the
source data of the descriptor update operation is taken from the
[VkWriteDescriptorSetInlineUniformBlock](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetInlineUniformBlock.html)
structure in the `pNext` chain of `VkWriteDescriptorSet`. When updating
descriptors with a `descriptorType` of
`VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR`, none of the
`pImageInfo`, `pBufferInfo`, or `pTexelBufferView` members are accessed,
instead the source data of the descriptor update operation is taken from
the
[VkWriteDescriptorSetAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetAccelerationStructureKHR.html)
structure in the `pNext` chain of `VkWriteDescriptorSet`. When updating
descriptors with a `descriptorType` of
`VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV`, none of the
`pImageInfo`, `pBufferInfo`, or `pTexelBufferView` members are accessed,
instead the source data of the descriptor update operation is taken from
the
[VkWriteDescriptorSetAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetAccelerationStructureNV.html)
structure in the `pNext` chain of `VkWriteDescriptorSet`.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDescriptorGetInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorGetInfoEXT.html),
[VkDescriptorPoolSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolSize.html),
[VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutBinding.html),
[VkDescriptorUpdateTemplateEntry](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorUpdateTemplateEntry.html),
[VkImageViewHandleInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewHandleInfoNVX.html),
[VkMutableDescriptorTypeListEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeListEXT.html),
[VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorType"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
