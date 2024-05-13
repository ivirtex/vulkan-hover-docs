# VkDescriptorPoolSize(3) Manual Page

## Name

VkDescriptorPoolSize - Structure specifying descriptor pool size



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDescriptorPoolSize` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkDescriptorPoolSize {
    VkDescriptorType    type;
    uint32_t            descriptorCount;
} VkDescriptorPoolSize;
```

## <a href="#_members" class="anchor"></a>Members

- `type` is the type of descriptor.

- `descriptorCount` is the number of descriptors of that type to
  allocate. If `type` is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` then
  `descriptorCount` is the number of bytes to allocate for descriptors
  of this type.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>When creating a descriptor pool that will contain descriptors for
combined image samplers of multi-planar formats, an application needs to
account for non-trivial descriptor consumption when choosing the
<code>descriptorCount</code> value, as indicated by <a
href="VkSamplerYcbcrConversionImageFormatProperties.html">VkSamplerYcbcrConversionImageFormatProperties</a>::<code>combinedImageSamplerDescriptorCount</code>.</p>
<p>For simplicity the application <strong>can</strong> use the <a
href="VkPhysicalDeviceMaintenance6PropertiesKHR.html">VkPhysicalDeviceMaintenance6PropertiesKHR</a>::<code>maxCombinedImageSamplerDescriptorCount</code>
property, which is sized to accommodate any and all <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion"
target="_blank" rel="noopener">formats that require a sampler
Y′C<sub>B</sub>C<sub>R</sub> conversion</a> supported by the
implementation.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkDescriptorPoolSize-descriptorCount-00302"
  id="VUID-VkDescriptorPoolSize-descriptorCount-00302"></a>
  VUID-VkDescriptorPoolSize-descriptorCount-00302  
  `descriptorCount` **must** be greater than `0`

- <a href="#VUID-VkDescriptorPoolSize-type-02218"
  id="VUID-VkDescriptorPoolSize-type-02218"></a>
  VUID-VkDescriptorPoolSize-type-02218  
  If `type` is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` then
  `descriptorCount` **must** be a multiple of `4`

Valid Usage (Implicit)

- <a href="#VUID-VkDescriptorPoolSize-type-parameter"
  id="VUID-VkDescriptorPoolSize-type-parameter"></a>
  VUID-VkDescriptorPoolSize-type-parameter  
  `type` **must** be a valid [VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorPoolCreateInfo.html),
[VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDescriptorPoolSize"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
