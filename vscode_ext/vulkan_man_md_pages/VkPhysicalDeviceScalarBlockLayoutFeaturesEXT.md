# VkPhysicalDeviceScalarBlockLayoutFeatures(3) Manual Page

## Name

VkPhysicalDeviceScalarBlockLayoutFeatures - Structure indicating support
for scalar block layouts



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceScalarBlockLayoutFeatures` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceScalarBlockLayoutFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           scalarBlockLayout;
} VkPhysicalDeviceScalarBlockLayoutFeatures;
```

or the equivalent

``` c
// Provided by VK_EXT_scalar_block_layout
typedef VkPhysicalDeviceScalarBlockLayoutFeatures VkPhysicalDeviceScalarBlockLayoutFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-scalarBlockLayout"></span>
  `scalarBlockLayout` indicates that the implementation supports the
  layout of resource blocks in shaders using <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-alignment-requirements"
  target="_blank" rel="noopener">scalar alignment</a>.

If the `VkPhysicalDeviceScalarBlockLayoutFeatures` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceScalarBlockLayoutFeatures` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceScalarBlockLayoutFeatures-sType-sType"
  id="VUID-VkPhysicalDeviceScalarBlockLayoutFeatures-sType-sType"></a>
  VUID-VkPhysicalDeviceScalarBlockLayoutFeatures-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SCALAR_BLOCK_LAYOUT_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_scalar_block_layout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_scalar_block_layout.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceScalarBlockLayoutFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
