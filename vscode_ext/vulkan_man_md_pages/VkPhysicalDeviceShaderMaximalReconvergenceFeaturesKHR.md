# VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR(3) Manual Page

## Name

VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR - Structure
describing support for shader maximal reconvergence by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR` structure is
defined as:

``` c
// Provided by VK_KHR_shader_maximal_reconvergence
typedef struct VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           shaderMaximalReconvergence;
} VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-shaderMaximalReconvergence"></span>
  `shaderMaximalReconvergence` specifies whether the implementation
  supports the shader execution mode `MaximallyReconvergesKHR`

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDevicePrivateDataFeaturesEXT` structure is included in
the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDevicePrivateDataFeaturesEXT` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR-sType-sType"
  id="VUID-VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR-sType-sType"></a>
  VUID-VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_MAXIMAL_RECONVERGENCE_FEATURES_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_shader_maximal_reconvergence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_maximal_reconvergence.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
