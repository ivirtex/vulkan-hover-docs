# VkDescriptorPoolSize(3) Manual Page

## Name

VkDescriptorPoolSize - Structure specifying descriptor pool size



## [](#_c_specification)C Specification

The `VkDescriptorPoolSize` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkDescriptorPoolSize {
    VkDescriptorType    type;
    uint32_t            descriptorCount;
} VkDescriptorPoolSize;
```

## [](#_members)Members

- `type` is the type of descriptor.
- `descriptorCount` is the number of descriptors of that type to allocate. If `type` is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` then `descriptorCount` is the number of bytes to allocate for descriptors of this type.

## [](#_description)Description

Note

When creating a descriptor pool that will contain descriptors for combined image samplers of [multi-planar formats](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-multiplanar), an application needs to account for non-trivial descriptor consumption when choosing the `descriptorCount` value, as indicated by [VkSamplerYcbcrConversionImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionImageFormatProperties.html)::`combinedImageSamplerDescriptorCount`.

For simplicity the application **can** use the [VkPhysicalDeviceMaintenance6Properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceMaintenance6Properties.html)::`maxCombinedImageSamplerDescriptorCount` property, which is sized to accommodate any and all [formats that require a sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-requiring-sampler-ycbcr-conversion) supported by the implementation.

Valid Usage

- [](#VUID-VkDescriptorPoolSize-descriptorCount-00302)VUID-VkDescriptorPoolSize-descriptorCount-00302  
  `descriptorCount` **must** be greater than `0`
- [](#VUID-VkDescriptorPoolSize-type-02218)VUID-VkDescriptorPoolSize-type-02218  
  If `type` is `VK_DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK` then `descriptorCount` **must** be a multiple of `4`

Valid Usage (Implicit)

- [](#VUID-VkDescriptorPoolSize-type-parameter)VUID-VkDescriptorPoolSize-type-parameter  
  `type` **must** be a valid [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html) value

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkDescriptorPoolCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorPoolCreateInfo.html), [VkDescriptorType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDescriptorType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDescriptorPoolSize).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0