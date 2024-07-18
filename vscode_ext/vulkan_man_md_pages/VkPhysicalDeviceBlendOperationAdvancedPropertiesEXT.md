# VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT - Structure
describing advanced blending limits that can be supported by an
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_blend_operation_advanced
typedef struct VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           advancedBlendMaxColorAttachments;
    VkBool32           advancedBlendIndependentBlend;
    VkBool32           advancedBlendNonPremultipliedSrcColor;
    VkBool32           advancedBlendNonPremultipliedDstColor;
    VkBool32           advancedBlendCorrelatedOverlap;
    VkBool32           advancedBlendAllOperations;
} VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-advancedBlendMaxColorAttachments"></span>
  `advancedBlendMaxColorAttachments` is one greater than the highest
  color attachment index that **can** be used in a subpass, for a
  pipeline that uses an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-blend-advanced"
  target="_blank" rel="noopener">advanced blend operation</a>.

- <span id="limits-advancedBlendIndependentBlend"></span>
  `advancedBlendIndependentBlend` specifies whether advanced blend
  operations **can** vary per-attachment.

- <span id="limits-advancedBlendNonPremultipliedSrcColor"></span>
  `advancedBlendNonPremultipliedSrcColor` specifies whether the source
  color **can** be treated as non-premultiplied. If this is `VK_FALSE`,
  then
  [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)::`srcPremultiplied`
  **must** be `VK_TRUE`.

- <span id="limits-advancedBlendNonPremultipliedDstColor"></span>
  `advancedBlendNonPremultipliedDstColor` specifies whether the
  destination color **can** be treated as non-premultiplied. If this is
  `VK_FALSE`, then
  [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)::`dstPremultiplied`
  **must** be `VK_TRUE`.

- <span id="limits-advancedBlendCorrelatedOverlap"></span>
  `advancedBlendCorrelatedOverlap` specifies whether the overlap mode
  **can** be treated as correlated. If this is `VK_FALSE`, then
  [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)::`blendOverlap`
  **must** be `VK_BLEND_OVERLAP_UNCORRELATED_EXT`.

- <span id="limits-advancedBlendAllOperations"></span>
  `advancedBlendAllOperations` specifies whether all advanced blend
  operation enums are supported. See the valid usage of
  [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendAttachmentState.html).

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_PROPERTIES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_blend_operation_advanced](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_blend_operation_advanced.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
