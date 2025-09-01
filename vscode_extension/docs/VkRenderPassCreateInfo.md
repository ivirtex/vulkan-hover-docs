# VkRenderPassCreateInfo(3) Manual Page

## Name

VkRenderPassCreateInfo - Structure specifying parameters of a newly created render pass



## [](#_c_specification)C Specification

The `VkRenderPassCreateInfo` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.2](#versions-1.2). See [Deprecated Functionality](#deprecation-renderpass2) for more information.

```c++
// Provided by VK_VERSION_1_0
typedef struct VkRenderPassCreateInfo {
    VkStructureType                   sType;
    const void*                       pNext;
    VkRenderPassCreateFlags           flags;
    uint32_t                          attachmentCount;
    const VkAttachmentDescription*    pAttachments;
    uint32_t                          subpassCount;
    const VkSubpassDescription*       pSubpasses;
    uint32_t                          dependencyCount;
    const VkSubpassDependency*        pDependencies;
} VkRenderPassCreateInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is a bitmask of [VkRenderPassCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateFlagBits.html)
- `attachmentCount` is the number of attachments used by this render pass.
- `pAttachments` is a pointer to an array of `attachmentCount` [VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescription.html) structures describing the attachments used by the render pass.
- `subpassCount` is the number of subpasses to create.
- `pSubpasses` is a pointer to an array of `subpassCount` [VkSubpassDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription.html) structures describing each subpass.
- `dependencyCount` is the number of memory dependencies between pairs of subpasses.
- `pDependencies` is a pointer to an array of `dependencyCount` [VkSubpassDependency](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDependency.html) structures describing dependencies between pairs of subpasses.

## [](#_description)Description

Note

Care should be taken to avoid a data race here; if any subpasses access attachments with overlapping memory locations, and one of those accesses is a write, a subpass dependency needs to be included between them.

Valid Usage

- [](#VUID-VkRenderPassCreateInfo-attachment-00834)VUID-VkRenderPassCreateInfo-attachment-00834  
  If the `attachment` member of any element of `pInputAttachments`, `pColorAttachments`, `pResolveAttachments` or `pDepthStencilAttachment`, or any element of `pPreserveAttachments` in any element of `pSubpasses` is not `VK_ATTACHMENT_UNUSED`, then it **must** be less than `attachmentCount`
- [](#VUID-VkRenderPassCreateInfo-fragmentDensityMapAttachment-06471)VUID-VkRenderPassCreateInfo-fragmentDensityMapAttachment-06471  
  If the pNext chain includes a [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html) structure and the `fragmentDensityMapAttachment` member is not `VK_ATTACHMENT_UNUSED`, then `attachment` **must** be less than `attachmentCount`
- [](#VUID-VkRenderPassCreateInfo-fragmentDensityMapLayered-10828)VUID-VkRenderPassCreateInfo-fragmentDensityMapLayered-10828  
  If the [`fragmentDensityMapLayered`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-fragmentDensityMapLayered) feature is not enabled, `flags` **must** not contain `VK_RENDER_PASS_CREATE_PER_LAYER_FRAGMENT_DENSITY_BIT_VALVE`
- [](#VUID-VkRenderPassCreateInfo-pAttachments-00836)VUID-VkRenderPassCreateInfo-pAttachments-00836  
  For any member of `pAttachments` with a `loadOp` equal to `VK_ATTACHMENT_LOAD_OP_CLEAR`, the first use of that attachment **must** not specify a `layout` equal to `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkRenderPassCreateInfo-pAttachments-02511)VUID-VkRenderPassCreateInfo-pAttachments-02511  
  For any member of `pAttachments` with a `stencilLoadOp` equal to `VK_ATTACHMENT_LOAD_OP_CLEAR`, the first use of that attachment **must** not specify a `layout` equal to `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL` or `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkRenderPassCreateInfo-pAttachments-01566)VUID-VkRenderPassCreateInfo-pAttachments-01566  
  For any member of `pAttachments` with a `loadOp` equal to `VK_ATTACHMENT_LOAD_OP_CLEAR`, the first use of that attachment **must** not specify a `layout` equal to `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`
- [](#VUID-VkRenderPassCreateInfo-pAttachments-01567)VUID-VkRenderPassCreateInfo-pAttachments-01567  
  For any member of `pAttachments` with a `stencilLoadOp` equal to `VK_ATTACHMENT_LOAD_OP_CLEAR`, the first use of that attachment **must** not specify a `layout` equal to `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`
- [](#VUID-VkRenderPassCreateInfo-pNext-01926)VUID-VkRenderPassCreateInfo-pNext-01926  
  If the `pNext` chain includes a [VkRenderPassInputAttachmentAspectCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassInputAttachmentAspectCreateInfo.html) structure, the `subpass` member of each element of its `pAspectReferences` member **must** be less than `subpassCount`
- [](#VUID-VkRenderPassCreateInfo-pNext-01927)VUID-VkRenderPassCreateInfo-pNext-01927  
  If the `pNext` chain includes a [VkRenderPassInputAttachmentAspectCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassInputAttachmentAspectCreateInfo.html) structure, the `inputAttachmentIndex` member of each element of its `pAspectReferences` member **must** be less than the value of `inputAttachmentCount` in the element of `pSubpasses` identified by its `subpass` member
- [](#VUID-VkRenderPassCreateInfo-pNext-01963)VUID-VkRenderPassCreateInfo-pNext-01963  
  If the `pNext` chain includes a [VkRenderPassInputAttachmentAspectCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassInputAttachmentAspectCreateInfo.html) structure, for any element of the `pInputAttachments` member of any element of `pSubpasses` where the `attachment` member is not `VK_ATTACHMENT_UNUSED`, the `aspectMask` member of the corresponding element of [VkRenderPassInputAttachmentAspectCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassInputAttachmentAspectCreateInfo.html)::`pAspectReferences` **must** only include aspects that are present in images of the format specified by the element of `pAttachments` at `attachment`
- [](#VUID-VkRenderPassCreateInfo-pNext-01928)VUID-VkRenderPassCreateInfo-pNext-01928  
  If the `pNext` chain includes a [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html) structure, and its `subpassCount` member is not zero, that member **must** be equal to the value of `subpassCount`
- [](#VUID-VkRenderPassCreateInfo-pNext-01929)VUID-VkRenderPassCreateInfo-pNext-01929  
  If the `pNext` chain includes a [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html) structure, if its `dependencyCount` member is not zero, it **must** be equal to `dependencyCount`
- [](#VUID-VkRenderPassCreateInfo-pNext-01930)VUID-VkRenderPassCreateInfo-pNext-01930  
  If the `pNext` chain includes a [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html) structure, for each non-zero element of `pViewOffsets`, the `srcSubpass` and `dstSubpass` members of `pDependencies` at the same index **must** not be equal
- [](#VUID-VkRenderPassCreateInfo-pNext-02512)VUID-VkRenderPassCreateInfo-pNext-02512  
  If the `pNext` chain includes a [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html) structure, for any element of `pDependencies` with a `dependencyFlags` member that does not include `VK_DEPENDENCY_VIEW_LOCAL_BIT`, the corresponding element of the `pViewOffsets` member of that [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html) instance **must** be `0`
- [](#VUID-VkRenderPassCreateInfo-pNext-02513)VUID-VkRenderPassCreateInfo-pNext-02513  
  If the `pNext` chain includes a [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html) structure, elements of its `pViewMasks` member **must** either all be `0`, or all not be `0`
- [](#VUID-VkRenderPassCreateInfo-pNext-02514)VUID-VkRenderPassCreateInfo-pNext-02514  
  If the `pNext` chain includes a [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html) structure, and each element of its `pViewMasks` member is `0`, the `dependencyFlags` member of each element of `pDependencies` **must** not include `VK_DEPENDENCY_VIEW_LOCAL_BIT`
- [](#VUID-VkRenderPassCreateInfo-pNext-02515)VUID-VkRenderPassCreateInfo-pNext-02515  
  If the `pNext` chain includes a [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html) structure, and each element of its `pViewMasks` member is `0`, its `correlationMaskCount` member **must** be `0`
- [](#VUID-VkRenderPassCreateInfo-pDependencies-00837)VUID-VkRenderPassCreateInfo-pDependencies-00837  
  For any element of `pDependencies`, if the `srcSubpass` is not `VK_SUBPASS_EXTERNAL`, all stage flags included in the `srcStageMask` member of that dependency **must** be `VK_PIPELINE_STAGE_ALL_COMMANDS_BIT` or a pipeline stage supported by the [pipeline](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-types) identified by the `pipelineBindPoint` member of the source subpass
- [](#VUID-VkRenderPassCreateInfo-pDependencies-00838)VUID-VkRenderPassCreateInfo-pDependencies-00838  
  For any element of `pDependencies`, if the `dstSubpass` is not `VK_SUBPASS_EXTERNAL`, all stage flags included in the `dstStageMask` member of that dependency **must** be `VK_PIPELINE_STAGE_ALL_COMMANDS_BIT` or a pipeline stage supported by the [pipeline](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages-types) identified by the `pipelineBindPoint` member of the destination subpass
- [](#VUID-VkRenderPassCreateInfo-pDependencies-06866)VUID-VkRenderPassCreateInfo-pDependencies-06866  
  For any element of `pDependencies`, if its `srcSubpass` is not `VK_SUBPASS_EXTERNAL`, it **must** be less than `subpassCount`
- [](#VUID-VkRenderPassCreateInfo-pDependencies-06867)VUID-VkRenderPassCreateInfo-pDependencies-06867  
  For any element of `pDependencies`, if its `dstSubpass` is not `VK_SUBPASS_EXTERNAL`, it **must** be less than `subpassCount`
- [](#VUID-VkRenderPassCreateInfo-pResolveAttachments-10647)VUID-VkRenderPassCreateInfo-pResolveAttachments-10647  
  If any element of `pResolveAttachments` of any element of `pSubpasses` references an attachment description with a format of `VK_FORMAT_UNDEFINED`, `VK_TILE_SHADING_RENDER_PASS_ENABLE_BIT_QCOM` **must** not be included in [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html)::`flags`
- [](#VUID-VkRenderPassCreateInfo-fragmentDensityMapAttachment-10648)VUID-VkRenderPassCreateInfo-fragmentDensityMapAttachment-10648  
  If [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html)::`fragmentDensityMapAttachment` is not `VK_ATTACHMENT_UNUSED`, `VK_TILE_SHADING_RENDER_PASS_ENABLE_BIT_QCOM` **must** not be included in [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html)::`flags`
- [](#VUID-VkRenderPassCreateInfo-None-10915)VUID-VkRenderPassCreateInfo-None-10915  
  If any subpass preserves an attachment, there **must** be a subpass dependency from a prior subpass which uses or preserves that attachment

Valid Usage (Implicit)

- [](#VUID-VkRenderPassCreateInfo-sType-sType)VUID-VkRenderPassCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO`
- [](#VUID-VkRenderPassCreateInfo-pNext-pNext)VUID-VkRenderPassCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html), [VkRenderPassInputAttachmentAspectCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassInputAttachmentAspectCreateInfo.html), [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassMultiviewCreateInfo.html), [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html), or [VkTileMemorySizeInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileMemorySizeInfoQCOM.html)
- [](#VUID-VkRenderPassCreateInfo-sType-unique)VUID-VkRenderPassCreateInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkRenderPassCreateInfo-flags-parameter)VUID-VkRenderPassCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of [VkRenderPassCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateFlagBits.html) values
- [](#VUID-VkRenderPassCreateInfo-pAttachments-parameter)VUID-VkRenderPassCreateInfo-pAttachments-parameter  
  If `attachmentCount` is not `0`, `pAttachments` **must** be a valid pointer to an array of `attachmentCount` valid [VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescription.html) structures
- [](#VUID-VkRenderPassCreateInfo-pSubpasses-parameter)VUID-VkRenderPassCreateInfo-pSubpasses-parameter  
  `pSubpasses` **must** be a valid pointer to an array of `subpassCount` valid [VkSubpassDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription.html) structures
- [](#VUID-VkRenderPassCreateInfo-pDependencies-parameter)VUID-VkRenderPassCreateInfo-pDependencies-parameter  
  If `dependencyCount` is not `0`, `pDependencies` **must** be a valid pointer to an array of `dependencyCount` valid [VkSubpassDependency](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDependency.html) structures
- [](#VUID-VkRenderPassCreateInfo-subpassCount-arraylength)VUID-VkRenderPassCreateInfo-subpassCount-arraylength  
  `subpassCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentDescription.html), [VkRenderPassCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSubpassDependency](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDependency.html), [VkSubpassDescription](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription.html), [vkCreateRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateRenderPass.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassCreateInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0