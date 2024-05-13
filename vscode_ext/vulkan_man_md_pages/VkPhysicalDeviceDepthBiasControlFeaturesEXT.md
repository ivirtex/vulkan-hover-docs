# VkPhysicalDeviceDepthBiasControlFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceDepthBiasControlFeaturesEXT - Structure indicating
support for depth bias scaling and representation control



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceDepthBiasControlFeaturesEXT` structure is defined
as:

``` c
// Provided by VK_EXT_depth_bias_control
typedef struct VkPhysicalDeviceDepthBiasControlFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           depthBiasControl;
    VkBool32           leastRepresentableValueForceUnormRepresentation;
    VkBool32           floatRepresentation;
    VkBool32           depthBiasExact;
} VkPhysicalDeviceDepthBiasControlFeaturesEXT;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following feature:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="features-depthBiasControl"></span> `depthBiasControl`
  indicates whether the implementation supports the
  `vkCmdSetDepthBias2EXT` command and the
  `VkDepthBiasRepresentationInfoEXT` structure.

- <span id="features-leastRepresentableValueForceUnormRepresentation"></span>
  `leastRepresentableValueForceUnormRepresentation` indicates whether
  the implementation supports using the
  `VK_DEPTH_BIAS_REPRESENTATION_LEAST_REPRESENTABLE_VALUE_FORCE_UNORM_EXT`
  depth bias representation.

- <span id="features-floatRepresentation"></span> `floatRepresentation`
  indicates whether the implementation supports using the
  `VK_DEPTH_BIAS_REPRESENTATION_FLOAT_EXT` depth bias representation.

- <span id="features-depthBiasExact"></span> `depthBiasExact` indicates
  whether the implementation supports forcing depth bias to not be
  scaled to ensure a minimum resolvable difference using
  `VkDepthBiasRepresentationInfoEXT`::`depthBiasExact`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceDepthBiasControlFeaturesEXT-sType-sType"
  id="VUID-VkPhysicalDeviceDepthBiasControlFeaturesEXT-sType-sType"></a>
  VUID-VkPhysicalDeviceDepthBiasControlFeaturesEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_DEPTH_BIAS_CONTROL_FEATURES_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_depth_bias_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_depth_bias_control.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceDepthBiasControlFeaturesEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
