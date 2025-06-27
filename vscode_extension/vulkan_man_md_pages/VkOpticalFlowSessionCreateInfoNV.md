# VkOpticalFlowSessionCreateInfoNV(3) Manual Page

## Name

VkOpticalFlowSessionCreateInfoNV - Structure specifying parameters of a
newly created optical flow session



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionCreateInfoNV.html)
structure is defined as:

``` c
// Provided by VK_NV_optical_flow
typedef struct VkOpticalFlowSessionCreateInfoNV {
    VkStructureType                      sType;
    void*                                pNext;
    uint32_t                             width;
    uint32_t                             height;
    VkFormat                             imageFormat;
    VkFormat                             flowVectorFormat;
    VkFormat                             costFormat;
    VkOpticalFlowGridSizeFlagsNV         outputGridSize;
    VkOpticalFlowGridSizeFlagsNV         hintGridSize;
    VkOpticalFlowPerformanceLevelNV      performanceLevel;
    VkOpticalFlowSessionCreateFlagsNV    flags;
} VkOpticalFlowSessionCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `width` is the width in pixels of the input or reference frame to be
  bound to this optical flow session.

- `height` is the height in pixels of the input or reference frame to be
  bound to this optical flow session.

- `imageFormat` is the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of the input and
  reference frame to be bound to this optical flow session.

- `flowVectorFormat` is the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of the flow vector
  maps (output or hint) to be bound to this optical flow session.

- `costFormat` is the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of the cost maps to be
  bound to this optical flow session.

- `outputGridSize` is exactly one bit of
  [VkOpticalFlowGridSizeFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowGridSizeFlagsNV.html)
  specifying the grid size of the output flow and cost maps to be bound
  to this optical flow session. The size of the output flow and cost
  maps is determined by `VkOpticalFlowSessionCreateInfoNV`::`width` and
  `VkOpticalFlowSessionCreateInfoNV`::`height` divided by
  `VkOpticalFlowSessionCreateInfoNV`::`outputGridSize`.

- `hintGridSize` is one exactly bit of
  [VkOpticalFlowGridSizeFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowGridSizeFlagsNV.html)
  specifying the grid size of the hint flow vector maps to be bound to
  this optical flow session. The size of the hint maps is determined by
  `VkOpticalFlowSessionCreateInfoNV`::`width` and
  `VkOpticalFlowSessionCreateInfoNV`::`height` divided by
  `VkOpticalFlowSessionCreateInfoNV`::`hintGridSize`.

- `performanceLevel` is the
  [VkOpticalFlowPerformanceLevelNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowPerformanceLevelNV.html)
  used for this optical flow session.

- `flags` are the
  [VkOpticalFlowSessionCreateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionCreateFlagsNV.html)
  used for this optical flow session.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-width-07581"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-width-07581"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-width-07581  
  `width` **must** be greater than or equal to
  `VkPhysicalDeviceOpticalFlowPropertiesNV`::`minWidth` and less than or
  equal to `VkPhysicalDeviceOpticalFlowPropertiesNV`::`maxWidth`

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-height-07582"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-height-07582"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-height-07582  
  `height` **must** be greater than or equal to
  `VkPhysicalDeviceOpticalFlowPropertiesNV`::`minHeight` and less than
  or equal to `VkPhysicalDeviceOpticalFlowPropertiesNV`::`maxHeight`

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-imageFormat-07583"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-imageFormat-07583"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-imageFormat-07583  
  `imageFormat` **must** be one of the formats returned by
  [vkGetPhysicalDeviceOpticalFlowImageFormatsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceOpticalFlowImageFormatsNV.html)
  for `VK_OPTICAL_FLOW_USAGE_INPUT_BIT_NV`

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-flowVectorFormat-07584"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-flowVectorFormat-07584"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-flowVectorFormat-07584  
  `flowVectorFormat` **must** be one of the formats returned by
  [vkGetPhysicalDeviceOpticalFlowImageFormatsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceOpticalFlowImageFormatsNV.html)
  for `VK_OPTICAL_FLOW_USAGE_OUTPUT_BIT_NV`

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-costFormat-07585"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-costFormat-07585"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-costFormat-07585  
  `costFormat` **must** be one of the formats returned by
  [vkGetPhysicalDeviceOpticalFlowImageFormatsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceOpticalFlowImageFormatsNV.html)
  for `VK_OPTICAL_FLOW_USAGE_COST_BIT_NV` if
  `VK_OPTICAL_FLOW_SESSION_CREATE_ENABLE_COST_BIT_NV` is set in `flags`

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-outputGridSize-07586"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-outputGridSize-07586"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-outputGridSize-07586  
  `outputGridSize` **must** be exactly one of the bits reported in
  `VkPhysicalDeviceOpticalFlowPropertiesNV`::`supportedOutputGridSizes`

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-hintGridSize-07587"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-hintGridSize-07587"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-hintGridSize-07587  
  `hintGridSize` **must** be exactly one of the bits reported in
  `VkPhysicalDeviceOpticalFlowPropertiesNV`::`supportedHintGridSizes` if
  `VK_OPTICAL_FLOW_SESSION_CREATE_ENABLE_HINT_BIT_NV` is set in `flags`

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-flags-07588"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-flags-07588"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-flags-07588  
  `VK_OPTICAL_FLOW_SESSION_CREATE_ENABLE_HINT_BIT_NV` **must** not be
  set in `flags` if
  `VkPhysicalDeviceOpticalFlowPropertiesNV`::`hintSupported` is
  `VK_FALSE`

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-flags-07589"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-flags-07589"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-flags-07589  
  `VK_OPTICAL_FLOW_SESSION_CREATE_ENABLE_COST_BIT_NV` **must** not be
  set in `flags` if
  `VkPhysicalDeviceOpticalFlowPropertiesNV`::`costSupported` is
  `VK_FALSE`

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-flags-07590"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-flags-07590"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-flags-07590  
  `VK_OPTICAL_FLOW_SESSION_CREATE_ENABLE_GLOBAL_FLOW_BIT_NV` **must**
  not be set in `flags` if
  `VkPhysicalDeviceOpticalFlowPropertiesNV`::`globalFlowSupported` is
  `VK_FALSE`

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-flags-07591"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-flags-07591"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-flags-07591  
  `VK_OPTICAL_FLOW_SESSION_CREATE_ALLOW_REGIONS_BIT_NV` **must** not be
  set in `flags` if
  `VkPhysicalDeviceOpticalFlowPropertiesNV`::`maxNumRegionsOfInterest`
  is 0

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-flags-07592"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-flags-07592"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-flags-07592  
  `VK_OPTICAL_FLOW_SESSION_CREATE_BOTH_DIRECTIONS_BIT_NV` **must** not
  be set in `flags` if
  `VkPhysicalDeviceOpticalFlowPropertiesNV`::`bidirectionalFlowSupported`
  is `VK_FALSE`

Valid Usage (Implicit)

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-sType-sType"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-sType-sType"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_OPTICAL_FLOW_SESSION_CREATE_INFO_NV`

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-pNext-pNext"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-pNext-pNext"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-pNext-pNext  
  `pNext` **must** be `NULL` or a pointer to a valid instance of
  [VkOpticalFlowSessionCreatePrivateDataInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionCreatePrivateDataInfoNV.html)

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-sType-unique"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-sType-unique"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-imageFormat-parameter"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-imageFormat-parameter"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-imageFormat-parameter  
  `imageFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a
  href="#VUID-VkOpticalFlowSessionCreateInfoNV-flowVectorFormat-parameter"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-flowVectorFormat-parameter"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-flowVectorFormat-parameter  
  `flowVectorFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-costFormat-parameter"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-costFormat-parameter"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-costFormat-parameter  
  If `costFormat` is not `0`, `costFormat` **must** be a valid
  [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a
  href="#VUID-VkOpticalFlowSessionCreateInfoNV-outputGridSize-parameter"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-outputGridSize-parameter"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-outputGridSize-parameter  
  `outputGridSize` **must** be a valid combination of
  [VkOpticalFlowGridSizeFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowGridSizeFlagBitsNV.html)
  values

- <a
  href="#VUID-VkOpticalFlowSessionCreateInfoNV-outputGridSize-requiredbitmask"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-outputGridSize-requiredbitmask"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-outputGridSize-requiredbitmask  
  `outputGridSize` **must** not be `0`

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-hintGridSize-parameter"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-hintGridSize-parameter"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-hintGridSize-parameter  
  `hintGridSize` **must** be a valid combination of
  [VkOpticalFlowGridSizeFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowGridSizeFlagBitsNV.html)
  values

- <a
  href="#VUID-VkOpticalFlowSessionCreateInfoNV-performanceLevel-parameter"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-performanceLevel-parameter"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-performanceLevel-parameter  
  If `performanceLevel` is not `0`, `performanceLevel` **must** be a
  valid
  [VkOpticalFlowPerformanceLevelNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowPerformanceLevelNV.html)
  value

- <a href="#VUID-VkOpticalFlowSessionCreateInfoNV-flags-parameter"
  id="VUID-VkOpticalFlowSessionCreateInfoNV-flags-parameter"></a>
  VUID-VkOpticalFlowSessionCreateInfoNV-flags-parameter  
  `flags` **must** be a valid combination of
  [VkOpticalFlowSessionCreateFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionCreateFlagBitsNV.html)
  values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_optical_flow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_optical_flow.html),
[VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkOpticalFlowGridSizeFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowGridSizeFlagsNV.html),
[VkOpticalFlowPerformanceLevelNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowPerformanceLevelNV.html),
[VkOpticalFlowSessionCreateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionCreateFlagsNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateOpticalFlowSessionNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkOpticalFlowSessionCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
