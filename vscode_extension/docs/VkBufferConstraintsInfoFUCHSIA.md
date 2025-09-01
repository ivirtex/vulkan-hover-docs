# VkBufferConstraintsInfoFUCHSIA(3) Manual Page

## Name

VkBufferConstraintsInfoFUCHSIA - Structure buffer-based buffer collection constraints



## [](#_c_specification)C Specification

The `VkBufferConstraintsInfoFUCHSIA` structure is defined as:

```c++
// Provided by VK_FUCHSIA_buffer_collection
typedef struct VkBufferConstraintsInfoFUCHSIA {
    VkStructureType                             sType;
    const void*                                 pNext;
    VkBufferCreateInfo                          createInfo;
    VkFormatFeatureFlags                        requiredFormatFeatures;
    VkBufferCollectionConstraintsInfoFUCHSIA    bufferCollectionConstraints;
} VkBufferConstraintsInfoFUCHSIA;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure
- `createInfo` a pointer to a [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html) struct describing the buffer attributes for the buffer collection
- `requiredFormatFeatures` bitmask of `VkFormatFeatureFlagBits` required features of the buffers in the buffer collection
- `bufferCollectionConstraints` is used to supply parameters for the negotiation and allocation of the buffer collection

## [](#_description)Description

Valid Usage

- [](#VUID-VkBufferConstraintsInfoFUCHSIA-requiredFormatFeatures-06404)VUID-VkBufferConstraintsInfoFUCHSIA-requiredFormatFeatures-06404  
  The `requiredFormatFeatures` bitmask of `VkFormatFeatureFlagBits` **must** be chosen from among the buffer compatible format features listed in [buffer compatible format features](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#buffer-compatible-format-features)

Valid Usage (Implicit)

- [](#VUID-VkBufferConstraintsInfoFUCHSIA-sType-sType)VUID-VkBufferConstraintsInfoFUCHSIA-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_CONSTRAINTS_INFO_FUCHSIA`
- [](#VUID-VkBufferConstraintsInfoFUCHSIA-pNext-pNext)VUID-VkBufferConstraintsInfoFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkBufferConstraintsInfoFUCHSIA-createInfo-parameter)VUID-VkBufferConstraintsInfoFUCHSIA-createInfo-parameter  
  `createInfo` **must** be a valid [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html) structure
- [](#VUID-VkBufferConstraintsInfoFUCHSIA-requiredFormatFeatures-parameter)VUID-VkBufferConstraintsInfoFUCHSIA-requiredFormatFeatures-parameter  
  `requiredFormatFeatures` **must** be a valid combination of [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits.html) values
- [](#VUID-VkBufferConstraintsInfoFUCHSIA-bufferCollectionConstraints-parameter)VUID-VkBufferConstraintsInfoFUCHSIA-bufferCollectionConstraints-parameter  
  `bufferCollectionConstraints` **must** be a valid [VkBufferCollectionConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionConstraintsInfoFUCHSIA.html) structure

## [](#_see_also)See Also

[VK\_FUCHSIA\_buffer\_collection](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_buffer_collection.html), [VkBufferCollectionConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionConstraintsInfoFUCHSIA.html), [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html), [VkFormatFeatureFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkSetBufferCollectionBufferConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionBufferConstraintsFUCHSIA.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferConstraintsInfoFUCHSIA).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0