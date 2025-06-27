# VkPhysicalDeviceOpticalFlowPropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceOpticalFlowPropertiesNV - Structure describing
properties supported by VK_NV_optical_flow



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceOpticalFlowPropertiesNV` structure is defined as:

``` c
// Provided by VK_NV_optical_flow
typedef struct VkPhysicalDeviceOpticalFlowPropertiesNV {
    VkStructureType                 sType;
    void*                           pNext;
    VkOpticalFlowGridSizeFlagsNV    supportedOutputGridSizes;
    VkOpticalFlowGridSizeFlagsNV    supportedHintGridSizes;
    VkBool32                        hintSupported;
    VkBool32                        costSupported;
    VkBool32                        bidirectionalFlowSupported;
    VkBool32                        globalFlowSupported;
    uint32_t                        minWidth;
    uint32_t                        minHeight;
    uint32_t                        maxWidth;
    uint32_t                        maxHeight;
    uint32_t                        maxNumRegionsOfInterest;
} VkPhysicalDeviceOpticalFlowPropertiesNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-supportedOutputGridSizes"></span>
  `supportedOutputGridSizes` are the supported
  [VkOpticalFlowGridSizeFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowGridSizeFlagsNV.html)
  which can be specified in
  `VkOpticalFlowSessionCreateInfoNV`::`outputGridSize`.

- <span id="limits-supportedHintGridSizes"></span>
  `supportedHintGridSizes` are the supported
  [VkOpticalFlowGridSizeFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowGridSizeFlagsNV.html)
  which can be specified in
  `VkOpticalFlowSessionCreateInfoNV`::`hintGridSize`.

- <span id="limits-hintSupported"></span> `hintSupported` is a boolean
  describing whether using hint flow vector map is supported in an
  optical flow session.

- <span id="limits-costSupported"></span> `costSupported` is a boolean
  describing whether cost map generation is supported in an optical flow
  session.

- <span id="limits-bidirectionalFlowSupported"></span>
  `bidirectionalFlowSupported` is a boolean describing whether
  bi-directional flow generation is supported in an optical flow
  session.

- <span id="limits-globalFlowSupported"></span> `globalFlowSupported` is
  a boolean describing whether global flow vector map generation is
  supported in an optical flow session.

- <span id="limits-minWidth"></span> `minWidth` is the minimum width in
  pixels for images used in an optical flow session.

- <span id="limits-minHeight"></span> `minHeight` is the minimum height
  in pixels for images used in an optical flow session.

- <span id="limits-maxWidth"></span> `maxWidth` is the maximum width in
  pixels for images used in an optical flow session.

- <span id="limits-maxHeight"></span> `maxHeight` is the maximum height
  in pixels for images used in an optical flow session.

- <span id="limits-maxNumRegionsOfInterest"></span>
  `maxNumRegionsOfInterest` is the maximum number of regions of interest
  which can be used in an optical flow session. If this
  `maxNumRegionsOfInterest` is 0, regions of interest are not supported
  in an optical flow session.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceOpticalFlowPropertiesNV` structure is included
in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceOpticalFlowPropertiesNV-sType-sType"
  id="VUID-VkPhysicalDeviceOpticalFlowPropertiesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceOpticalFlowPropertiesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_OPTICAL_FLOW_PROPERTIES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_optical_flow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_optical_flow.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkOpticalFlowGridSizeFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowGridSizeFlagsNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceOpticalFlowPropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
