# VkPipelineColorWriteCreateInfoEXT(3) Manual Page

## Name

VkPipelineColorWriteCreateInfoEXT - Structure specifying color write
state of a newly created pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineColorWriteCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_color_write_enable
typedef struct VkPipelineColorWriteCreateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           attachmentCount;
    const VkBool32*    pColorWriteEnables;
} VkPipelineColorWriteCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `attachmentCount` is the number of [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html) elements
  in `pColorWriteEnables`.

- `pColorWriteEnables` is a pointer to an array of per target attachment
  boolean values specifying whether color writes are enabled for the
  given attachment.

## <a href="#_description" class="anchor"></a>Description

When this structure is included in the `pNext` chain of
[VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html),
it defines per-attachment color write state. If this structure is not
included in the `pNext` chain, it is equivalent to specifying this
structure with `attachmentCount` equal to the `attachmentCount` member
of
[VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html),
and `pColorWriteEnables` pointing to an array of as many `VK_TRUE`
values.

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-colorWriteEnable"
target="_blank" rel="noopener"><code>colorWriteEnable</code></a> feature
is not enabled on the device, all [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html) elements in
the `pColorWriteEnables` array **must** be `VK_TRUE`.

Color Write Enable interacts with the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#framebuffer-color-write-mask"
target="_blank" rel="noopener">Color Write Mask</a> as follows:

- If `colorWriteEnable` is `VK_TRUE`, writes to the attachment are
  determined by the `colorWriteMask`.

- If `colorWriteEnable` is `VK_FALSE`, the `colorWriteMask` is ignored
  and writes to all components of the attachment are disabled. This is
  equivalent to specifying a `colorWriteMask` of 0.

Valid Usage

- <a href="#VUID-VkPipelineColorWriteCreateInfoEXT-pAttachments-04801"
  id="VUID-VkPipelineColorWriteCreateInfoEXT-pAttachments-04801"></a>
  VUID-VkPipelineColorWriteCreateInfoEXT-pAttachments-04801  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-colorWriteEnable"
  target="_blank" rel="noopener"><code>colorWriteEnable</code></a>
  feature is not enabled, all elements of `pColorWriteEnables` **must**
  be `VK_TRUE`

- <a href="#VUID-VkPipelineColorWriteCreateInfoEXT-attachmentCount-07608"
  id="VUID-VkPipelineColorWriteCreateInfoEXT-attachmentCount-07608"></a>
  VUID-VkPipelineColorWriteCreateInfoEXT-attachmentCount-07608  
  If the pipeline is being created with
  `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT`,
  `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT`,
  `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT`, or
  `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT` dynamic states not set,
  `attachmentCount` **must** be equal to the `attachmentCount` member of
  the `VkPipelineColorBlendStateCreateInfo` structure specified during
  pipeline creation

- <a href="#VUID-VkPipelineColorWriteCreateInfoEXT-attachmentCount-06655"
  id="VUID-VkPipelineColorWriteCreateInfoEXT-attachmentCount-06655"></a>
  VUID-VkPipelineColorWriteCreateInfoEXT-attachmentCount-06655  
  `attachmentCount` **must** be less than or equal to the
  `maxColorAttachments` member of `VkPhysicalDeviceLimits`

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineColorWriteCreateInfoEXT-sType-sType"
  id="VUID-VkPipelineColorWriteCreateInfoEXT-sType-sType"></a>
  VUID-VkPipelineColorWriteCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_COLOR_WRITE_CREATE_INFO_EXT`

- <a
  href="#VUID-VkPipelineColorWriteCreateInfoEXT-pColorWriteEnables-parameter"
  id="VUID-VkPipelineColorWriteCreateInfoEXT-pColorWriteEnables-parameter"></a>
  VUID-VkPipelineColorWriteCreateInfoEXT-pColorWriteEnables-parameter  
  If `attachmentCount` is not `0`, `pColorWriteEnables` **must** be a
  valid pointer to an array of `attachmentCount`
  [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_color_write_enable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_color_write_enable.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineColorWriteCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
