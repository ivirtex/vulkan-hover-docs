# VkBufferConstraintsInfoFUCHSIA(3) Manual Page

## Name

VkBufferConstraintsInfoFUCHSIA - Structure buffer-based buffer
collection constraints



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBufferConstraintsInfoFUCHSIA` structure is defined as:

``` c
// Provided by VK_FUCHSIA_buffer_collection
typedef struct VkBufferConstraintsInfoFUCHSIA {
    VkStructureType                             sType;
    const void*                                 pNext;
    VkBufferCreateInfo                          createInfo;
    VkFormatFeatureFlags                        requiredFormatFeatures;
    VkBufferCollectionConstraintsInfoFUCHSIA    bufferCollectionConstraints;
} VkBufferConstraintsInfoFUCHSIA;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this structure

- `pBufferCreateInfo` a pointer to a
  [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) struct describing the
  buffer attributes for the buffer collection

- `requiredFormatFeatures` bitmask of `VkFormatFeatureFlagBits` required
  features of the buffers in the buffer collection

- `bufferCollectionConstraints` is used to supply parameters for the
  negotiation and allocation of the buffer collection

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkBufferConstraintsInfoFUCHSIA-requiredFormatFeatures-06404"
  id="VUID-VkBufferConstraintsInfoFUCHSIA-requiredFormatFeatures-06404"></a>
  VUID-VkBufferConstraintsInfoFUCHSIA-requiredFormatFeatures-06404  
  The `requiredFormatFeatures` bitmask of `VkFormatFeatureFlagBits`
  **must** be chosen from among the buffer compatible format features
  listed in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#buffer-compatible-format-features"
  target="_blank" rel="noopener">buffer compatible format features</a>

Valid Usage (Implicit)

- <a href="#VUID-VkBufferConstraintsInfoFUCHSIA-sType-sType"
  id="VUID-VkBufferConstraintsInfoFUCHSIA-sType-sType"></a>
  VUID-VkBufferConstraintsInfoFUCHSIA-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BUFFER_CONSTRAINTS_INFO_FUCHSIA`

- <a href="#VUID-VkBufferConstraintsInfoFUCHSIA-pNext-pNext"
  id="VUID-VkBufferConstraintsInfoFUCHSIA-pNext-pNext"></a>
  VUID-VkBufferConstraintsInfoFUCHSIA-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkBufferConstraintsInfoFUCHSIA-createInfo-parameter"
  id="VUID-VkBufferConstraintsInfoFUCHSIA-createInfo-parameter"></a>
  VUID-VkBufferConstraintsInfoFUCHSIA-createInfo-parameter  
  `createInfo` **must** be a valid
  [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) structure

- <a
  href="#VUID-VkBufferConstraintsInfoFUCHSIA-requiredFormatFeatures-parameter"
  id="VUID-VkBufferConstraintsInfoFUCHSIA-requiredFormatFeatures-parameter"></a>
  VUID-VkBufferConstraintsInfoFUCHSIA-requiredFormatFeatures-parameter  
  `requiredFormatFeatures` **must** be a valid combination of
  [VkFormatFeatureFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits.html) values

- <a
  href="#VUID-VkBufferConstraintsInfoFUCHSIA-bufferCollectionConstraints-parameter"
  id="VUID-VkBufferConstraintsInfoFUCHSIA-bufferCollectionConstraints-parameter"></a>
  VUID-VkBufferConstraintsInfoFUCHSIA-bufferCollectionConstraints-parameter  
  `bufferCollectionConstraints` **must** be a valid
  [VkBufferCollectionConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionConstraintsInfoFUCHSIA.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_FUCHSIA_buffer_collection](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_buffer_collection.html),
[VkBufferCollectionConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionConstraintsInfoFUCHSIA.html),
[VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html),
[VkFormatFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkSetBufferCollectionBufferConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetBufferCollectionBufferConstraintsFUCHSIA.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferConstraintsInfoFUCHSIA"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
