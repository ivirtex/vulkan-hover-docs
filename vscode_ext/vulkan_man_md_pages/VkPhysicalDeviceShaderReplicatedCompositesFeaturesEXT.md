# VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT - Structure
describing whether support for replicated composites in SPIR-V is
enabled



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT` structure is
defined as:

``` c
// Provided by VK_EXT_shader_replicated_composites
typedef struct VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderReplicatedComposites;
} VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-shaderReplicatedComposites"></span>
  `shaderReplicatedComposites` specifies whether shader modules **can**
  declare the `ReplicatedCompositesEXT` capability.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT` structure
is included in the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT` **can** also be
used in the `pNext` chain of
[VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to selectively enable
these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_REPLICATED_COMPOSITES_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_shader_replicated_composites](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_shader_replicated_composites.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderReplicatedCompositesFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
