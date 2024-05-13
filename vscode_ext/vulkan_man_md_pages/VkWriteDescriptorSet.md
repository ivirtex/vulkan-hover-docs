# VkWriteDescriptorSet(3) Manual Page

## Name

VkWriteDescriptorSet - Structure specifying the parameters of a
descriptor set write operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkWriteDescriptorSet` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `dstSet` is the destination descriptor set to update.

- `dstBinding` is the descriptor binding within that set.

- `dstArrayElement` is the starting element in that array. If the
  descriptor binding identified by `dstSet` and `dstBinding` has a
  descriptor type of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` then
  `dstArrayElement` specifies the starting byte offset within the
  binding.

- `descriptorCount` is the number of descriptors to update. If the
  descriptor binding identified by `dstSet` and `dstBinding` has a
  descriptor type of `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, then
  `descriptorCount` specifies the number of bytes to update. Otherwise,
  `descriptorCount` is one of

  - the number of elements in `pImageInfo`

  - the number of elements in `pBufferInfo`

  - the number of elements in `pTexelBufferView`

  - a value matching the `dataSize` member of a
    [VkWriteDescriptorSetInlineUniformBlock](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetInlineUniformBlock.html)
    structure in the `pNext` chain

  - a value matching the `accelerationStructureCount` of a
    [VkWriteDescriptorSetAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetAccelerationStructureKHR.html)
    structure in the `pNext` chain

- `descriptorType` is a [VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html)
  specifying the type of each descriptor in `pImageInfo`, `pBufferInfo`,
  or `pTexelBufferView`, as described below. If
  `VkDescriptorSetLayoutBinding` for `dstSet` at `dstBinding` is not
  equal to `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, `descriptorType` **must**
  be the same type as the `descriptorType` specified in
  `VkDescriptorSetLayoutBinding` for `dstSet` at `dstBinding`. The type
  of the descriptor also controls which array the descriptors are taken
  from.

- `pImageInfo` is a pointer to an array of
  [VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorImageInfo.html) structures or is
  ignored, as described below.

- `pBufferInfo` is a pointer to an array of
  [VkDescriptorBufferInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferInfo.html) structures or is
  ignored, as described below.

- `pTexelBufferView` is a pointer to an array of
  [VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html) handles as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-buffer-views"
  target="_blank" rel="noopener">Buffer Views</a> section or is ignored,
  as described below.

## <a href="#_description" class="anchor"></a>Description

Only one of `pImageInfo`, `pBufferInfo`, or `pTexelBufferView` members
is used according to the descriptor type specified in the
`descriptorType` member of the containing `VkWriteDescriptorSet`
structure, or none of them in case `descriptorType` is
`VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, in which case the source data
for the descriptor writes is taken from the
[VkWriteDescriptorSetInlineUniformBlock](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetInlineUniformBlock.html)
structure included in the `pNext` chain of `VkWriteDescriptorSet`, or if
`descriptorType` is `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR`, in
which case the source data for the descriptor writes is taken from the
[VkWriteDescriptorSetAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetAccelerationStructureKHR.html)
structure in the `pNext` chain of `VkWriteDescriptorSet`, or if
`descriptorType` is `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV`, in
which case the source data for the descriptor writes is taken from the
[VkWriteDescriptorSetAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetAccelerationStructureNV.html)
structure in the `pNext` chain of `VkWriteDescriptorSet`, as specified
below.

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nullDescriptor"
target="_blank" rel="noopener"><code>nullDescriptor</code></a> feature
is enabled, the buffer, acceleration structure, imageView, or bufferView
**can** be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html). Loads from a null
descriptor return zero values and stores and atomics to a null
descriptor are discarded. A null acceleration structure descriptor
results in the miss shader being invoked.

If the destination descriptor is a mutable descriptor, the active
descriptor type for the destination descriptor becomes `descriptorType`.

If the `dstBinding` has fewer than `descriptorCount` array elements
remaining starting from `dstArrayElement`, then the remainder will be
used to update the subsequent binding - `dstBinding`+1 starting at array
element zero. If a binding has a `descriptorCount` of zero, it is
skipped. This behavior applies recursively, with the update affecting
consecutive bindings as needed to update all `descriptorCount`
descriptors. Consecutive bindings **must** have identical
[VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html),
[VkShaderStageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStageFlags.html),
[VkDescriptorBindingFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlagBits.html), and
immutable samplers references. In addition, if the
[VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html) is
`VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, the supported descriptor types in
[VkMutableDescriptorTypeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMutableDescriptorTypeCreateInfoEXT.html)
**must** be equally defined.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The same behavior applies to bindings with a descriptor type of
<code>VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK</code> where
<code>descriptorCount</code> specifies the number of bytes to update
while <code>dstArrayElement</code> specifies the starting byte offset,
thus in this case if the <code>dstBinding</code> has a smaller byte size
than the sum of <code>dstArrayElement</code> and
<code>descriptorCount</code>, then the remainder will be used to update
the subsequent binding - <code>dstBinding</code>+1 starting at offset
zero. This falls out as a special case of the above rule.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkWriteDescriptorSet-dstBinding-00315"
  id="VUID-VkWriteDescriptorSet-dstBinding-00315"></a>
  VUID-VkWriteDescriptorSet-dstBinding-00315  
  `dstBinding` **must** be less than or equal to the maximum value of
  `binding` of all
  [VkDescriptorSetLayoutBinding](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutBinding.html)
  structures specified when `dstSet`’s descriptor set layout was created

- <a href="#VUID-VkWriteDescriptorSet-dstBinding-00316"
  id="VUID-VkWriteDescriptorSet-dstBinding-00316"></a>
  VUID-VkWriteDescriptorSet-dstBinding-00316  
  `dstBinding` **must** be a binding with a non-zero `descriptorCount`

- <a href="#VUID-VkWriteDescriptorSet-descriptorCount-00317"
  id="VUID-VkWriteDescriptorSet-descriptorCount-00317"></a>
  VUID-VkWriteDescriptorSet-descriptorCount-00317  
  All consecutive bindings updated via a single `VkWriteDescriptorSet`
  structure, except those with a `descriptorCount` of zero, **must**
  have identical `descriptorType` and `stageFlags`

- <a href="#VUID-VkWriteDescriptorSet-descriptorCount-00318"
  id="VUID-VkWriteDescriptorSet-descriptorCount-00318"></a>
  VUID-VkWriteDescriptorSet-descriptorCount-00318  
  All consecutive bindings updated via a single `VkWriteDescriptorSet`
  structure, except those with a `descriptorCount` of zero, **must** all
  either use immutable samplers or **must** all not use immutable
  samplers

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-00319"
  id="VUID-VkWriteDescriptorSet-descriptorType-00319"></a>
  VUID-VkWriteDescriptorSet-descriptorType-00319  
  `descriptorType` **must** match the type of `dstBinding` within
  `dstSet`

- <a href="#VUID-VkWriteDescriptorSet-dstSet-00320"
  id="VUID-VkWriteDescriptorSet-dstSet-00320"></a>
  VUID-VkWriteDescriptorSet-dstSet-00320  
  `dstSet` **must** be a valid [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html)
  handle

- <a href="#VUID-VkWriteDescriptorSet-dstArrayElement-00321"
  id="VUID-VkWriteDescriptorSet-dstArrayElement-00321"></a>
  VUID-VkWriteDescriptorSet-dstArrayElement-00321  
  The sum of `dstArrayElement` and `descriptorCount` **must** be less
  than or equal to the number of array elements in the descriptor set
  binding specified by `dstBinding`, and all applicable consecutive
  bindings, as described by <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-updates-consecutive"
  class="bare" target="_blank"
  rel="noopener">https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-updates-consecutive</a>

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-02219"
  id="VUID-VkWriteDescriptorSet-descriptorType-02219"></a>
  VUID-VkWriteDescriptorSet-descriptorType-02219  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`,
  `dstArrayElement` **must** be an integer multiple of `4`

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-02220"
  id="VUID-VkWriteDescriptorSet-descriptorType-02220"></a>
  VUID-VkWriteDescriptorSet-descriptorType-02220  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`,
  `descriptorCount` **must** be an integer multiple of `4`

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-02994"
  id="VUID-VkWriteDescriptorSet-descriptorType-02994"></a>
  VUID-VkWriteDescriptorSet-descriptorType-02994  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` or
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`, each element of
  `pTexelBufferView` **must** be either a valid `VkBufferView` handle or
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-02995"
  id="VUID-VkWriteDescriptorSet-descriptorType-02995"></a>
  VUID-VkWriteDescriptorSet-descriptorType-02995  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER` or
  `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER` and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nullDescriptor"
  target="_blank" rel="noopener"><code>nullDescriptor</code></a> feature
  is not enabled, each element of `pTexelBufferView` **must** not be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-00324"
  id="VUID-VkWriteDescriptorSet-descriptorType-00324"></a>
  VUID-VkWriteDescriptorSet-descriptorType-00324  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`,
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`,
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`, or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`, `pBufferInfo` **must** be
  a valid pointer to an array of `descriptorCount` valid
  `VkDescriptorBufferInfo` structures

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-00325"
  id="VUID-VkWriteDescriptorSet-descriptorType-00325"></a>
  VUID-VkWriteDescriptorSet-descriptorType-00325  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLER` or
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, and `dstSet` was not
  allocated with a layout that included immutable samplers for
  `dstBinding` with `descriptorType`, the `sampler` member of each
  element of `pImageInfo` **must** be a valid `VkSampler` object

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-02996"
  id="VUID-VkWriteDescriptorSet-descriptorType-02996"></a>
  VUID-VkWriteDescriptorSet-descriptorType-02996  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, or
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, the `imageView` member of each
  element of `pImageInfo` **must** be either a valid `VkImageView`
  handle or [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-02997"
  id="VUID-VkWriteDescriptorSet-descriptorType-02997"></a>
  VUID-VkWriteDescriptorSet-descriptorType-02997  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, or
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, and the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nullDescriptor"
  target="_blank" rel="noopener"><code>nullDescriptor</code></a> feature
  is not enabled, the `imageView` member of each element of `pImageInfo`
  **must** not be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-07683"
  id="VUID-VkWriteDescriptorSet-descriptorType-07683"></a>
  VUID-VkWriteDescriptorSet-descriptorType-07683  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, the
  `imageView` member of each element of `pImageInfo` **must** not be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-02221"
  id="VUID-VkWriteDescriptorSet-descriptorType-02221"></a>
  VUID-VkWriteDescriptorSet-descriptorType-02221  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK`, the
  `pNext` chain **must** include a
  [VkWriteDescriptorSetInlineUniformBlock](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetInlineUniformBlock.html)
  structure whose `dataSize` member equals `descriptorCount`

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-02382"
  id="VUID-VkWriteDescriptorSet-descriptorType-02382"></a>
  VUID-VkWriteDescriptorSet-descriptorType-02382  
  If `descriptorType` is
  `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR`, the `pNext` chain
  **must** include a
  [VkWriteDescriptorSetAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetAccelerationStructureKHR.html)
  structure whose `accelerationStructureCount` member equals
  `descriptorCount`

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-03817"
  id="VUID-VkWriteDescriptorSet-descriptorType-03817"></a>
  VUID-VkWriteDescriptorSet-descriptorType-03817  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV`,
  the `pNext` chain **must** include a
  [VkWriteDescriptorSetAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetAccelerationStructureNV.html)
  structure whose `accelerationStructureCount` member equals
  `descriptorCount`

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-01946"
  id="VUID-VkWriteDescriptorSet-descriptorType-01946"></a>
  VUID-VkWriteDescriptorSet-descriptorType-01946  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`, then the
  `imageView` member of each `pImageInfo` element **must** have been
  created without a `VkSamplerYcbcrConversionInfo` structure in its
  `pNext` chain

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-02738"
  id="VUID-VkWriteDescriptorSet-descriptorType-02738"></a>
  VUID-VkWriteDescriptorSet-descriptorType-02738  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`,
  and if any element of `pImageInfo` has an `imageView` member that was
  created with a `VkSamplerYcbcrConversionInfo` structure in its `pNext`
  chain, then `dstSet` **must** have been allocated with a layout that
  included immutable samplers for `dstBinding`, and the corresponding
  immutable sampler **must** have been created with an *identically
  defined* `VkSamplerYcbcrConversionInfo` object

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-01948"
  id="VUID-VkWriteDescriptorSet-descriptorType-01948"></a>
  VUID-VkWriteDescriptorSet-descriptorType-01948  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`,
  and `dstSet` was allocated with a layout that included immutable
  samplers for `dstBinding`, then the `imageView` member of each element
  of `pImageInfo` which corresponds to an immutable sampler that enables
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-YCbCr-conversion"
  target="_blank" rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion</a> **must** have been created with a
  `VkSamplerYcbcrConversionInfo` structure in its `pNext` chain with an
  *identically defined* `VkSamplerYcbcrConversionInfo` to the
  corresponding immutable sampler

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-09506"
  id="VUID-VkWriteDescriptorSet-descriptorType-09506"></a>
  VUID-VkWriteDescriptorSet-descriptorType-09506  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`,
  `dstSet` was allocated with a layout that included immutable samplers
  for `dstBinding`, and those samplers enable <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#samplers-YCbCr-conversion"
  target="_blank" rel="noopener">sampler Y′C<sub>B</sub>C<sub>R</sub>
  conversion</a>, then `imageView` **must** not be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-00327"
  id="VUID-VkWriteDescriptorSet-descriptorType-00327"></a>
  VUID-VkWriteDescriptorSet-descriptorType-00327  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` or
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`, the `offset` member of
  each element of `pBufferInfo` **must** be a multiple of
  `VkPhysicalDeviceLimits`::`minUniformBufferOffsetAlignment`

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-00328"
  id="VUID-VkWriteDescriptorSet-descriptorType-00328"></a>
  VUID-VkWriteDescriptorSet-descriptorType-00328  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`, the `offset` member of
  each element of `pBufferInfo` **must** be a multiple of
  `VkPhysicalDeviceLimits`::`minStorageBufferOffsetAlignment`

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-00329"
  id="VUID-VkWriteDescriptorSet-descriptorType-00329"></a>
  VUID-VkWriteDescriptorSet-descriptorType-00329  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER`,
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`,
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER`, or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`, and the `buffer` member
  of any element of `pBufferInfo` is the handle of a non-sparse buffer,
  then that buffer **must** be bound completely and contiguously to a
  single `VkDeviceMemory` object

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-00330"
  id="VUID-VkWriteDescriptorSet-descriptorType-00330"></a>
  VUID-VkWriteDescriptorSet-descriptorType-00330  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` or
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`, the `buffer` member of
  each element of `pBufferInfo` **must** have been created with
  `VK_BUFFER_USAGE_UNIFORM_BUFFER_BIT` set

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-00331"
  id="VUID-VkWriteDescriptorSet-descriptorType-00331"></a>
  VUID-VkWriteDescriptorSet-descriptorType-00331  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`, the `buffer` member of
  each element of `pBufferInfo` **must** have been created with
  `VK_BUFFER_USAGE_STORAGE_BUFFER_BIT` set

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-00332"
  id="VUID-VkWriteDescriptorSet-descriptorType-00332"></a>
  VUID-VkWriteDescriptorSet-descriptorType-00332  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER` or
  `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC`, the `range` member of
  each element of `pBufferInfo`, or the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#buffer-info-effective-range"
  target="_blank" rel="noopener">effective range</a> if `range` is
  `VK_WHOLE_SIZE`, **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxUniformBufferRange`

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-00333"
  id="VUID-VkWriteDescriptorSet-descriptorType-00333"></a>
  VUID-VkWriteDescriptorSet-descriptorType-00333  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER` or
  `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC`, the `range` member of
  each element of `pBufferInfo`, or the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#buffer-info-effective-range"
  target="_blank" rel="noopener">effective range</a> if `range` is
  `VK_WHOLE_SIZE`, **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxStorageBufferRange`

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-08765"
  id="VUID-VkWriteDescriptorSet-descriptorType-08765"></a>
  VUID-VkWriteDescriptorSet-descriptorType-08765  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER`, the
  `pTexelBufferView` <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-buffer-views-usage"
  target="_blank" rel="noopener">buffer view usage</a> **must** include
  `VK_BUFFER_USAGE_UNIFORM_TEXEL_BUFFER_BIT`

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-08766"
  id="VUID-VkWriteDescriptorSet-descriptorType-08766"></a>
  VUID-VkWriteDescriptorSet-descriptorType-08766  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER`, the
  `pTexelBufferView` <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-buffer-views-usage"
  target="_blank" rel="noopener">buffer view usage</a> **must** include
  `VK_BUFFER_USAGE_STORAGE_TEXEL_BUFFER_BIT`

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-00336"
  id="VUID-VkWriteDescriptorSet-descriptorType-00336"></a>
  VUID-VkWriteDescriptorSet-descriptorType-00336  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` or
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, the `imageView` member of each
  element of `pImageInfo` **must** have been created with the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
  target="_blank" rel="noopener">identity swizzle</a>

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-00337"
  id="VUID-VkWriteDescriptorSet-descriptorType-00337"></a>
  VUID-VkWriteDescriptorSet-descriptorType-00337  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` or
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, the `imageView` member of
  each element of `pImageInfo` **must** have been created with
  `VK_IMAGE_USAGE_SAMPLED_BIT` set

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-04149"
  id="VUID-VkWriteDescriptorSet-descriptorType-04149"></a>
  VUID-VkWriteDescriptorSet-descriptorType-04149  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE` the
  `imageLayout` member of each element of `pImageInfo` **must** be a
  member of the list given in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-sampledimage"
  target="_blank" rel="noopener">Sampled Image</a>

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-04150"
  id="VUID-VkWriteDescriptorSet-descriptorType-04150"></a>
  VUID-VkWriteDescriptorSet-descriptorType-04150  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER` the
  `imageLayout` member of each element of `pImageInfo` **must** be a
  member of the list given in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-combinedimagesampler"
  target="_blank" rel="noopener">Combined Image Sampler</a>

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-04151"
  id="VUID-VkWriteDescriptorSet-descriptorType-04151"></a>
  VUID-VkWriteDescriptorSet-descriptorType-04151  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT` the
  `imageLayout` member of each element of `pImageInfo` **must** be a
  member of the list given in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-inputattachment"
  target="_blank" rel="noopener">Input Attachment</a>

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-04152"
  id="VUID-VkWriteDescriptorSet-descriptorType-04152"></a>
  VUID-VkWriteDescriptorSet-descriptorType-04152  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE` the
  `imageLayout` member of each element of `pImageInfo` **must** be a
  member of the list given in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-storageimage"
  target="_blank" rel="noopener">Storage Image</a>

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-00338"
  id="VUID-VkWriteDescriptorSet-descriptorType-00338"></a>
  VUID-VkWriteDescriptorSet-descriptorType-00338  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, the
  `imageView` member of each element of `pImageInfo` **must** have been
  created with `VK_IMAGE_USAGE_INPUT_ATTACHMENT_BIT` set

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-00339"
  id="VUID-VkWriteDescriptorSet-descriptorType-00339"></a>
  VUID-VkWriteDescriptorSet-descriptorType-00339  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, the
  `imageView` member of each element of `pImageInfo` **must** have been
  created with `VK_IMAGE_USAGE_STORAGE_BIT` set

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-02752"
  id="VUID-VkWriteDescriptorSet-descriptorType-02752"></a>
  VUID-VkWriteDescriptorSet-descriptorType-02752  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLER`, then `dstSet`
  **must** not have been allocated with a layout that included immutable
  samplers for `dstBinding`

- <a href="#VUID-VkWriteDescriptorSet-dstSet-04611"
  id="VUID-VkWriteDescriptorSet-dstSet-04611"></a>
  VUID-VkWriteDescriptorSet-dstSet-04611  
  If the `VkDescriptorSetLayoutBinding` for `dstSet` at `dstBinding` is
  `VK_DESCRIPTOR_TYPE_MUTABLE_EXT`, the new active descriptor type
  `descriptorType` **must** exist in the corresponding
  `pMutableDescriptorTypeLists` list for `dstBinding`

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-06450"
  id="VUID-VkWriteDescriptorSet-descriptorType-06450"></a>
  VUID-VkWriteDescriptorSet-descriptorType-06450  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`, the
  `imageView` member of each element of `pImageInfo` **must** have
  either been created without a
  [VkImageViewMinLodCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewMinLodCreateInfoEXT.html)
  included in the `pNext` chain or with a
  [VkImageViewMinLodCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewMinLodCreateInfoEXT.html)::`minLod`
  of `0.0`

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-06942"
  id="VUID-VkWriteDescriptorSet-descriptorType-06942"></a>
  VUID-VkWriteDescriptorSet-descriptorType-06942  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM`,
  the `imageView` member of each element of `pImageInfo` **must** have
  been created with a view created with an `image` created with
  `VK_IMAGE_USAGE_SAMPLE_WEIGHT_BIT_QCOM`

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-06943"
  id="VUID-VkWriteDescriptorSet-descriptorType-06943"></a>
  VUID-VkWriteDescriptorSet-descriptorType-06943  
  If `descriptorType` is `VK_DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM`,
  the `imageView` member of each element of `pImageInfo` **must** have
  been created with a view created with an `image` created with
  `VK_IMAGE_USAGE_SAMPLE_BLOCK_MATCH_BIT_QCOM`

Valid Usage (Implicit)

- <a href="#VUID-VkWriteDescriptorSet-sType-sType"
  id="VUID-VkWriteDescriptorSet-sType-sType"></a>
  VUID-VkWriteDescriptorSet-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_WRITE_DESCRIPTOR_SET`

- <a href="#VUID-VkWriteDescriptorSet-pNext-pNext"
  id="VUID-VkWriteDescriptorSet-pNext-pNext"></a>
  VUID-VkWriteDescriptorSet-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkWriteDescriptorSetAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetAccelerationStructureKHR.html),
  [VkWriteDescriptorSetAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetAccelerationStructureNV.html),
  or
  [VkWriteDescriptorSetInlineUniformBlock](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSetInlineUniformBlock.html)

- <a href="#VUID-VkWriteDescriptorSet-sType-unique"
  id="VUID-VkWriteDescriptorSet-sType-unique"></a>
  VUID-VkWriteDescriptorSet-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkWriteDescriptorSet-descriptorType-parameter"
  id="VUID-VkWriteDescriptorSet-descriptorType-parameter"></a>
  VUID-VkWriteDescriptorSet-descriptorType-parameter  
  `descriptorType` **must** be a valid
  [VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html) value

- <a href="#VUID-VkWriteDescriptorSet-descriptorCount-arraylength"
  id="VUID-VkWriteDescriptorSet-descriptorCount-arraylength"></a>
  VUID-VkWriteDescriptorSet-descriptorCount-arraylength  
  `descriptorCount` **must** be greater than `0`

- <a href="#VUID-VkWriteDescriptorSet-commonparent"
  id="VUID-VkWriteDescriptorSet-commonparent"></a>
  VUID-VkWriteDescriptorSet-commonparent  
  Both of `dstSet`, and the elements of `pTexelBufferView` that are
  valid handles of non-ignored parameters **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkBufferView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferView.html),
[VkDescriptorBufferInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferInfo.html),
[VkDescriptorImageInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorImageInfo.html),
[VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html),
[VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html),
[VkPushDescriptorSetInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushDescriptorSetInfoKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCmdPushDescriptorSetKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdPushDescriptorSetKHR.html),
[vkUpdateDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkUpdateDescriptorSets.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkWriteDescriptorSet"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
