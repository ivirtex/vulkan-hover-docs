# VkWriteDescriptorSet(3) Manual Page

## Name

VkWriteDescriptorSet - Structure specifying the parameters of a descriptor set write operation



## [](#_c_specification)C Specification

The `VkWriteDescriptorSet` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkWriteDescriptorSet {
    VkStructureType                  sType;
    const void*                      pNext;
    VkDescriptorSet                  dstSet;
    uint32_t                         dstBinding;
    uint32_t                         dstArrayElement;
    uint32_t                         descriptorCount;
    VkDescriptorType                 descriptorType;
    const VkDescriptorImageInfo*     pImageInfo;
    const VkDescriptorBufferInfo*    pBufferInfo;
    const VkBufferView*              pTexelBufferView;
} VkWriteDescriptorSet;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `dstSet` is the destination descriptor set to update.
- `dstBinding` is the descriptor binding within that set.
- `dstArrayElement` is the starting element in that array. If the descriptor binding identified by `dstSet` and `dstBinding` has a descriptor type of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` then `dstArrayElement` specifies the starting byte offset within the binding.
- `descriptorCount` is the number of descriptors to update. If the descriptor binding identified by `dstSet` and `dstBinding` has a descriptor type of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, then `descriptorCount` specifies the number of bytes to update. Otherwise, `descriptorCount` is one of
  
  - the number of elements in `pImageInfo`
  - the number of elements in `pBufferInfo`
  - the number of elements in `pTexelBufferView`
  - a value matching the `dataSize` member of a [VkWriteDescriptorSetInlineUniformBlock](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetInlineUniformBlock.html) structure in the `pNext` chain
  - a value matching the `accelerationStructureCount` of a [VkWriteDescriptorSetAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetAccelerationStructureKHR.html) or [VkWriteDescriptorSetAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetAccelerationStructureNV.html) structure in the `pNext` chain
  - a value matching the `descriptorCount` of a [VkWriteDescriptorSetTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetTensorARM.html) structure in the `pNext` chain
- `descriptorType` is a [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html) specifying the type of each descriptor in `pImageInfo`, `pBufferInfo`, or `pTexelBufferView`, as described below. If `VkDescriptorSetLayoutBinding` for `dstSet` at `dstBinding` is not equal to `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, `descriptorType` **must** be the same type as the `descriptorType` specified in `VkDescriptorSetLayoutBinding` for `dstSet` at `dstBinding`. The type of the descriptor also controls which array the descriptors are taken from.
- `pImageInfo` is a pointer to an array of [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorImageInfo.html) structures or is ignored, as described below.
- `pBufferInfo` is a pointer to an array of [VkDescriptorBufferInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferInfo.html) structures or is ignored, as described below.
- `pTexelBufferView` is a pointer to an array of [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html) handles as described in the [Buffer Views](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-buffer-views) section or is ignored, as described below.

## [](#_description)Description

Only one of `pImageInfo`, `pBufferInfo`, or `pTexelBufferView` members is used according to the descriptor type specified in the `descriptorType` member of the containing `VkWriteDescriptorSet` structure, or none of them in case `descriptorType` is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, in which case the source data for the descriptor writes is taken from the [VkWriteDescriptorSetInlineUniformBlock](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetInlineUniformBlock.html) structure included in the `pNext` chain of `VkWriteDescriptorSet`, or if `descriptorType` is `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR`, in which case the source data for the descriptor writes is taken from the [VkWriteDescriptorSetAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetAccelerationStructureKHR.html) structure in the `pNext` chain of `VkWriteDescriptorSet`, or if `descriptorType` is `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV`, in which case the source data for the descriptor writes is taken from the [VkWriteDescriptorSetAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetAccelerationStructureNV.html) structure in the `pNext` chain of `VkWriteDescriptorSet`, or if `descriptorType` is `VK_DESCRIPTOR_TYPE_TENSOR_ARM`, in which case the source data for the descriptor writes is taken from the instance of [VkWriteDescriptorSetTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetTensorARM.html) in the `pNext` chain of `VkWriteDescriptorSet`, as specified below.

If the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is enabled, the buffer, acceleration structure, tensor, imageView, or bufferView **can** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html). Loads from a null descriptor return zero values and stores and atomics to a null descriptor are discarded. A null acceleration structure descriptor results in the miss shader being invoked.

If the destination descriptor is a mutable descriptor, the active descriptor type for the destination descriptor becomes `descriptorType`.

Consecutive Binding Updates

If the `dstBinding` has fewer than `descriptorCount` array elements remaining starting from `dstArrayElement`, then the remainder will be used to update the subsequent binding - `dstBinding`+1 starting at array element zero. If a binding has a `descriptorCount` of zero, it is skipped. This behavior applies recursively, with the update affecting consecutive bindings as needed to update all `descriptorCount` descriptors. Consecutive bindings **must** have identical [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html), [VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderStageFlags.html), [VkDescriptorBindingFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBindingFlagBits.html), and immutable samplers references. In addition, if the [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html) is `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, the supported descriptor types in [VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMutableDescriptorTypeCreateInfoEXT.html) **must** be equally defined.

Note

The same behavior applies to bindings with a descriptor type of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` where `descriptorCount` specifies the number of bytes to update while `dstArrayElement` specifies the starting byte offset, thus in this case if the `dstBinding` has a smaller byte size than the sum of `dstArrayElement` and `descriptorCount`, then the remainder will be used to update the subsequent binding - `dstBinding`+1 starting at offset zero. This falls out as a special case of the above rule.

Valid Usage

- [](#VUID-VkWriteDescriptorSet-dstBinding-00315)VUID-VkWriteDescriptorSet-dstBinding-00315  
  `dstBinding` **must** be less than or equal to the maximum value of `binding` of all [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutBinding.html) structures specified when `dstSet`’s descriptor set layout was created
- [](#VUID-VkWriteDescriptorSet-dstBinding-00316)VUID-VkWriteDescriptorSet-dstBinding-00316  
  `dstBinding` **must** be a binding with a non-zero `descriptorCount`
- [](#VUID-VkWriteDescriptorSet-dstBinding-10009)VUID-VkWriteDescriptorSet-dstBinding-10009  
  `dstBinding` **must** be a binding with a non-zero [VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSetLayoutCreateInfo.html)::`bindingCount`
- [](#VUID-VkWriteDescriptorSet-descriptorCount-00317)VUID-VkWriteDescriptorSet-descriptorCount-00317  
  All consecutive bindings updated via a single `VkWriteDescriptorSet` structure, except those with a `descriptorCount` of zero, **must** have identical `descriptorType`
- [](#VUID-VkWriteDescriptorSet-descriptorCount-10776)VUID-VkWriteDescriptorSet-descriptorCount-10776  
  All consecutive bindings updated via a single `VkWriteDescriptorSet` structure, except those with a `descriptorCount` of zero, **must** have identical `stageFlags`
- [](#VUID-VkWriteDescriptorSet-descriptorCount-00318)VUID-VkWriteDescriptorSet-descriptorCount-00318  
  All consecutive bindings updated via a single `VkWriteDescriptorSet` structure, except those with a `descriptorCount` of zero, **must** all either use immutable samplers or **must** all not use immutable samplers
- [](#VUID-VkWriteDescriptorSet-descriptorCount-10777)VUID-VkWriteDescriptorSet-descriptorCount-10777  
  All consecutive bindings updated via a single `VkWriteDescriptorSet` structure, except those with a `descriptorCount` of zero, **must** have identical [VkDescriptorBindingFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBindingFlagBits.html)
- [](#VUID-VkWriteDescriptorSet-descriptorType-00319)VUID-VkWriteDescriptorSet-descriptorType-00319  
  `descriptorType` **must** match the type of `dstBinding` within `dstSet`
- [](#VUID-VkWriteDescriptorSet-dstSet-00320)VUID-VkWriteDescriptorSet-dstSet-00320  
  `dstSet` **must** be a valid [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html) handle
- [](#VUID-VkWriteDescriptorSet-dstArrayElement-00321)VUID-VkWriteDescriptorSet-dstArrayElement-00321  
  The sum of `dstArrayElement` and `descriptorCount` **must** be less than or equal to the number of array elements in the descriptor set binding specified by `dstBinding`, and all applicable [consecutive bindings](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-updates-consecutive)
- [](#VUID-VkWriteDescriptorSet-descriptorType-02219)VUID-VkWriteDescriptorSet-descriptorType-02219  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, `dstArrayElement` **must** be an integer multiple of `4`
- [](#VUID-VkWriteDescriptorSet-descriptorType-02220)VUID-VkWriteDescriptorSet-descriptorType-02220  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, `descriptorCount` **must** be an integer multiple of `4`
- [](#VUID-VkWriteDescriptorSet-descriptorType-02994)VUID-VkWriteDescriptorSet-descriptorType-02994  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` or `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`, each element of `pTexelBufferView` **must** be either a valid `VkBufferView` handle or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkWriteDescriptorSet-descriptorType-02995)VUID-VkWriteDescriptorSet-descriptorType-02995  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` or `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` and the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, each element of `pTexelBufferView` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkWriteDescriptorSet-descriptorType-00324)VUID-VkWriteDescriptorSet-descriptorType-00324  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`, `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`, `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`, or `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`, `pBufferInfo` **must** be a valid pointer to an array of `descriptorCount` valid `VkDescriptorBufferInfo` structures
- [](#VUID-VkWriteDescriptorSet-descriptorType-00325)VUID-VkWriteDescriptorSet-descriptorType-00325  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLER` or `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, and `dstSet` was not allocated with a layout that included immutable samplers for `dstBinding` with `descriptorType`, the `sampler` member of each element of `pImageInfo` **must** be a valid `VkSampler` object
- [](#VUID-VkWriteDescriptorSet-descriptorType-02996)VUID-VkWriteDescriptorSet-descriptorType-02996  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, or `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, the `imageView` member of each element of `pImageInfo` **must** be either a valid `VkImageView` handle or [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkWriteDescriptorSet-descriptorType-02997)VUID-VkWriteDescriptorSet-descriptorType-02997  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, or `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, and the [`nullDescriptor`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-nullDescriptor) feature is not enabled, the `imageView` member of each element of `pImageInfo` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkWriteDescriptorSet-descriptorType-07683)VUID-VkWriteDescriptorSet-descriptorType-07683  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, the `imageView` member of each element of `pImageInfo` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkWriteDescriptorSet-descriptorType-02221)VUID-VkWriteDescriptorSet-descriptorType-02221  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, the `pNext` chain **must** include a [VkWriteDescriptorSetInlineUniformBlock](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetInlineUniformBlock.html) structure whose `dataSize` member equals `descriptorCount`
- [](#VUID-VkWriteDescriptorSet-descriptorType-02382)VUID-VkWriteDescriptorSet-descriptorType-02382  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR`, the `pNext` chain **must** include a [VkWriteDescriptorSetAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetAccelerationStructureKHR.html) structure whose `accelerationStructureCount` member equals `descriptorCount`
- [](#VUID-VkWriteDescriptorSet-descriptorType-03817)VUID-VkWriteDescriptorSet-descriptorType-03817  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV`, the `pNext` chain **must** include a [VkWriteDescriptorSetAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetAccelerationStructureNV.html) structure whose `accelerationStructureCount` member equals `descriptorCount`
- [](#VUID-VkWriteDescriptorSet-descriptorType-09945)VUID-VkWriteDescriptorSet-descriptorType-09945  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_TENSOR_ARM`, the `pNext` chain **must** include a [VkWriteDescriptorSetTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetTensorARM.html) structure whose `tensorViewCount` member equals `descriptorCount`
- [](#VUID-VkWriteDescriptorSet-descriptorType-01946)VUID-VkWriteDescriptorSet-descriptorType-01946  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, then the `imageView` member of each `pImageInfo` element **must** have been created without a `VkSamplerYcbcrConversionInfo` structure in its `pNext` chain
- [](#VUID-VkWriteDescriptorSet-descriptorType-02738)VUID-VkWriteDescriptorSet-descriptorType-02738  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, and if any element of `pImageInfo` has an `imageView` member that was created with a `VkSamplerYcbcrConversionInfo` structure in its `pNext` chain, then `dstSet` **must** have been allocated with a layout that included immutable samplers for `dstBinding`, and the corresponding immutable sampler **must** have been created with an [identically defined](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-identically-defined) `VkSamplerYcbcrConversionInfo` object
- [](#VUID-VkWriteDescriptorSet-descriptorType-01948)VUID-VkWriteDescriptorSet-descriptorType-01948  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, and `dstSet` was allocated with a layout that included immutable samplers for `dstBinding`, then the `imageView` member of each element of `pImageInfo` which corresponds to an immutable sampler that enables [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion) **must** have been created with a `VkSamplerYcbcrConversionInfo` structure in its `pNext` chain with an [identically defined](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#glossary-identically-defined) `VkSamplerYcbcrConversionInfo` to the corresponding immutable sampler
- [](#VUID-VkWriteDescriptorSet-descriptorType-09506)VUID-VkWriteDescriptorSet-descriptorType-09506  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, `dstSet` was allocated with a layout that included immutable samplers for `dstBinding`, and those samplers enable [sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#samplers-YCbCr-conversion), then `imageView` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkWriteDescriptorSet-descriptorType-00327)VUID-VkWriteDescriptorSet-descriptorType-00327  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` or `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`, the `offset` member of each element of `pBufferInfo` **must** be a multiple of `VkPhysicalDeviceLimits`::`minUniformBufferOffsetAlignment`
- [](#VUID-VkWriteDescriptorSet-descriptorType-00328)VUID-VkWriteDescriptorSet-descriptorType-00328  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` or `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`, the `offset` member of each element of `pBufferInfo` **must** be a multiple of `VkPhysicalDeviceLimits`::`minStorageBufferOffsetAlignment`
- [](#VUID-VkWriteDescriptorSet-descriptorType-00329)VUID-VkWriteDescriptorSet-descriptorType-00329  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`, `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`, `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`, or `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`, and the `buffer` member of any element of `pBufferInfo` is the handle of a non-sparse buffer, then that buffer **must** be bound completely and contiguously to a single `VkDeviceMemory` object
- [](#VUID-VkWriteDescriptorSet-descriptorType-00330)VUID-VkWriteDescriptorSet-descriptorType-00330  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` or `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`, the `buffer` member of each element of `pBufferInfo` **must** have been created with `VK_BUFFER_USAGE_UNIFORM_BUFFER_BIT` set
- [](#VUID-VkWriteDescriptorSet-descriptorType-00331)VUID-VkWriteDescriptorSet-descriptorType-00331  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` or `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`, the `buffer` member of each element of `pBufferInfo` **must** have been created with `VK_BUFFER_USAGE_STORAGE_BUFFER_BIT` set
- [](#VUID-VkWriteDescriptorSet-descriptorType-00332)VUID-VkWriteDescriptorSet-descriptorType-00332  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` or `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`, the `range` member of each element of `pBufferInfo`, or the [effective range](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#buffer-info-effective-range) if `range` is `VK_WHOLE_SIZE`, **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxUniformBufferRange`
- [](#VUID-VkWriteDescriptorSet-descriptorType-00333)VUID-VkWriteDescriptorSet-descriptorType-00333  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` or `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`, the `range` member of each element of `pBufferInfo`, or the [effective range](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#buffer-info-effective-range) if `range` is `VK_WHOLE_SIZE`, **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxStorageBufferRange`
- [](#VUID-VkWriteDescriptorSet-descriptorType-08765)VUID-VkWriteDescriptorSet-descriptorType-08765  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`, the `pTexelBufferView` [buffer view usage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-buffer-views-usage) **must** include `VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT`
- [](#VUID-VkWriteDescriptorSet-descriptorType-08766)VUID-VkWriteDescriptorSet-descriptorType-08766  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`, the `pTexelBufferView` [buffer view usage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-buffer-views-usage) **must** include `VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT`
- [](#VUID-VkWriteDescriptorSet-descriptorType-00336)VUID-VkWriteDescriptorSet-descriptorType-00336  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` or `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, the `imageView` member of each element of `pImageInfo` **must** have been created with the [identity swizzle](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-image-views-identity-mappings)
- [](#VUID-VkWriteDescriptorSet-descriptorType-00337)VUID-VkWriteDescriptorSet-descriptorType-00337  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` or `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, the `imageView` member of each element of `pImageInfo` **must** have been created with `VK_IMAGE_USAGE_SAMPLED_BIT` set
- [](#VUID-VkWriteDescriptorSet-descriptorType-04149)VUID-VkWriteDescriptorSet-descriptorType-04149  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` the `imageLayout` member of each element of `pImageInfo` **must** be a member of the list given in [Sampled Image](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-sampledimage)
- [](#VUID-VkWriteDescriptorSet-descriptorType-04150)VUID-VkWriteDescriptorSet-descriptorType-04150  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` the `imageLayout` member of each element of `pImageInfo` **must** be a member of the list given in [Combined Image Sampler](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-combinedimagesampler)
- [](#VUID-VkWriteDescriptorSet-descriptorType-04151)VUID-VkWriteDescriptorSet-descriptorType-04151  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` the `imageLayout` member of each element of `pImageInfo` **must** be a member of the list given in [Input Attachment](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-inputattachment)
- [](#VUID-VkWriteDescriptorSet-descriptorType-04152)VUID-VkWriteDescriptorSet-descriptorType-04152  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` the `imageLayout` member of each element of `pImageInfo` **must** be a member of the list given in [Storage Image](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#descriptorsets-storageimage)
- [](#VUID-VkWriteDescriptorSet-descriptorType-00338)VUID-VkWriteDescriptorSet-descriptorType-00338  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, the `imageView` member of each element of `pImageInfo` **must** have been created with `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` set
- [](#VUID-VkWriteDescriptorSet-descriptorType-00339)VUID-VkWriteDescriptorSet-descriptorType-00339  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, the `imageView` member of each element of `pImageInfo` **must** have been created with `VK_IMAGE_USAGE_STORAGE_BIT` set
- [](#VUID-VkWriteDescriptorSet-descriptorType-02752)VUID-VkWriteDescriptorSet-descriptorType-02752  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLER`, then `dstSet` **must** not have been allocated with a layout that included immutable samplers for `dstBinding`
- [](#VUID-VkWriteDescriptorSet-dstSet-04611)VUID-VkWriteDescriptorSet-dstSet-04611  
  If the `VkDescriptorSetLayoutBinding` for `dstSet` at `dstBinding` is `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, the new active descriptor type `descriptorType` **must** exist in the corresponding `pMutableDescriptorTypeLists` list for `dstBinding`
- [](#VUID-VkWriteDescriptorSet-descriptorType-06450)VUID-VkWriteDescriptorSet-descriptorType-06450  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, the `imageView` member of each element of `pImageInfo` **must** have either been created without a [VkImageViewMinLodCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewMinLodCreateInfoEXT.html) included in the `pNext` chain or with a [VkImageViewMinLodCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageViewMinLodCreateInfoEXT.html)::`minLod` of `0.0`
- [](#VUID-VkWriteDescriptorSet-descriptorType-06942)VUID-VkWriteDescriptorSet-descriptorType-06942  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM`, the `imageView` member of each element of `pImageInfo` **must** have been created with a view created with an `image` created with `VK_IMAGE_USAGE_SAMPLE_WEIGHT_BIT_QCOM`
- [](#VUID-VkWriteDescriptorSet-descriptorType-06943)VUID-VkWriteDescriptorSet-descriptorType-06943  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM`, the `imageView` member of each element of `pImageInfo` **must** have been created with a view created with an `image` created with `VK_IMAGE_USAGE_SAMPLE_BLOCK_MATCH_BIT_QCOM`

Valid Usage (Implicit)

- [](#VUID-VkWriteDescriptorSet-sType-sType)VUID-VkWriteDescriptorSet-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET`
- [](#VUID-VkWriteDescriptorSet-pNext-pNext)VUID-VkWriteDescriptorSet-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkWriteDescriptorSetAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetAccelerationStructureKHR.html), [VkWriteDescriptorSetAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetAccelerationStructureNV.html), [VkWriteDescriptorSetInlineUniformBlock](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetInlineUniformBlock.html), [VkWriteDescriptorSetPartitionedAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetPartitionedAccelerationStructureNV.html), or [VkWriteDescriptorSetTensorARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkWriteDescriptorSetTensorARM.html)
- [](#VUID-VkWriteDescriptorSet-sType-unique)VUID-VkWriteDescriptorSet-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkWriteDescriptorSet-descriptorType-parameter)VUID-VkWriteDescriptorSet-descriptorType-parameter  
  `descriptorType` **must** be a valid [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html) value
- [](#VUID-VkWriteDescriptorSet-descriptorCount-arraylength)VUID-VkWriteDescriptorSet-descriptorCount-arraylength  
  `descriptorCount` **must** be greater than `0`
- [](#VUID-VkWriteDescriptorSet-commonparent)VUID-VkWriteDescriptorSet-commonparent  
  Both of `dstSet`, and the elements of `pTexelBufferView` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBufferView](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferView.html), [VkDescriptorBufferInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorBufferInfo.html), [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorImageInfo.html), [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorSet.html), [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html), [VkPushDescriptorSetInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPushDescriptorSetInfo.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdPushDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSet.html), [vkCmdPushDescriptorSet](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdPushDescriptorSet.html), [vkUpdateDescriptorSets](https://registry.khronos.org/vulkan/specs/latest/man/html/vkUpdateDescriptorSets.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkWriteDescriptorSet).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0