# VkPipelineRepresentativeFragmentTestStateCreateInfoNV(3) Manual Page

## Name

VkPipelineRepresentativeFragmentTestStateCreateInfoNV - Structure
specifying representative fragment test



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html)
includes a `VkPipelineRepresentativeFragmentTestStateCreateInfoNV`
structure, then that structure includes parameters controlling the
representative fragment test.

The `VkPipelineRepresentativeFragmentTestStateCreateInfoNV` structure is
defined as:

``` c
// Provided by VK_NV_representative_fragment_test
typedef struct VkPipelineRepresentativeFragmentTestStateCreateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           representativeFragmentTestEnable;
} VkPipelineRepresentativeFragmentTestStateCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `representativeFragmentTestEnable` controls whether the representative
  fragment test is enabled.

## <a href="#_description" class="anchor"></a>Description

If this structure is not included in the `pNext` chain,
`representativeFragmentTestEnable` is considered to be `VK_FALSE`, and
the representative fragment test is disabled.

If the active fragment shader does not specify the `EarlyFragmentTests`
execution mode, the representative fragment shader test has no effect,
even if enabled.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineRepresentativeFragmentTestStateCreateInfoNV-sType-sType"
  id="VUID-VkPipelineRepresentativeFragmentTestStateCreateInfoNV-sType-sType"></a>
  VUID-VkPipelineRepresentativeFragmentTestStateCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_REPRESENTATIVE_FRAGMENT_TEST_STATE_CREATE_INFO_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_representative_fragment_test](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_representative_fragment_test.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineRepresentativeFragmentTestStateCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
