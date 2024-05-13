# VkRenderPassCreateInfo(3) Manual Page

## Name

VkRenderPassCreateInfo - Structure specifying parameters of a newly
created render pass



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkRenderPassCreateInfo` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkRenderPassCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateFlagBits.html)

- `attachmentCount` is the number of attachments used by this render
  pass.

- `pAttachments` is a pointer to an array of `attachmentCount`
  [VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription.html) structures
  describing the attachments used by the render pass.

- `subpassCount` is the number of subpasses to create.

- `pSubpasses` is a pointer to an array of `subpassCount`
  [VkSubpassDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription.html) structures
  describing each subpass.

- `dependencyCount` is the number of memory dependencies between pairs
  of subpasses.

- `pDependencies` is a pointer to an array of `dependencyCount`
  [VkSubpassDependency](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDependency.html) structures describing
  dependencies between pairs of subpasses.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Care should be taken to avoid a data race here; if any subpasses
access attachments with overlapping memory locations, and one of those
accesses is a write, a subpass dependency needs to be included between
them.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkRenderPassCreateInfo-attachment-00834"
  id="VUID-VkRenderPassCreateInfo-attachment-00834"></a>
  VUID-VkRenderPassCreateInfo-attachment-00834  
  If the `attachment` member of any element of `pInputAttachments`,
  `pColorAttachments`, `pResolveAttachments` or
  `pDepthStencilAttachment`, or any element of `pPreserveAttachments` in
  any element of `pSubpasses` is not `VK_ATTACHMENT_UNUSED`, then it
  **must** be less than `attachmentCount`

- <a
  href="#VUID-VkRenderPassCreateInfo-fragmentDensityMapAttachment-06471"
  id="VUID-VkRenderPassCreateInfo-fragmentDensityMapAttachment-06471"></a>
  VUID-VkRenderPassCreateInfo-fragmentDensityMapAttachment-06471  
  If the pNext chain includes a
  [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html)
  structure and the `fragmentDensityMapAttachment` member is not
  `VK_ATTACHMENT_UNUSED`, then `attachment` **must** be less than
  `attachmentCount`

- <a href="#VUID-VkRenderPassCreateInfo-pAttachments-00836"
  id="VUID-VkRenderPassCreateInfo-pAttachments-00836"></a>
  VUID-VkRenderPassCreateInfo-pAttachments-00836  
  For any member of `pAttachments` with a `loadOp` equal to
  `VK_ATTACHMENT_LOAD_OP_CLEAR`, the first use of that attachment
  **must** not specify a `layout` equal to
  `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkRenderPassCreateInfo-pAttachments-02511"
  id="VUID-VkRenderPassCreateInfo-pAttachments-02511"></a>
  VUID-VkRenderPassCreateInfo-pAttachments-02511  
  For any member of `pAttachments` with a `stencilLoadOp` equal to
  `VK_ATTACHMENT_LOAD_OP_CLEAR`, the first use of that attachment
  **must** not specify a `layout` equal to
  `VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL` or
  `VK_IMAGE_LAYOUT_DEPTH_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkRenderPassCreateInfo-pAttachments-01566"
  id="VUID-VkRenderPassCreateInfo-pAttachments-01566"></a>
  VUID-VkRenderPassCreateInfo-pAttachments-01566  
  For any member of `pAttachments` with a `loadOp` equal to
  `VK_ATTACHMENT_LOAD_OP_CLEAR`, the first use of that attachment
  **must** not specify a `layout` equal to
  `VK_IMAGE_LAYOUT_DEPTH_READ_ONLY_STENCIL_ATTACHMENT_OPTIMAL`

- <a href="#VUID-VkRenderPassCreateInfo-pAttachments-01567"
  id="VUID-VkRenderPassCreateInfo-pAttachments-01567"></a>
  VUID-VkRenderPassCreateInfo-pAttachments-01567  
  For any member of `pAttachments` with a `stencilLoadOp` equal to
  `VK_ATTACHMENT_LOAD_OP_CLEAR`, the first use of that attachment
  **must** not specify a `layout` equal to
  `VK_IMAGE_LAYOUT_DEPTH_ATTACHMENT_STENCIL_READ_ONLY_OPTIMAL`

- <a href="#VUID-VkRenderPassCreateInfo-pNext-01926"
  id="VUID-VkRenderPassCreateInfo-pNext-01926"></a>
  VUID-VkRenderPassCreateInfo-pNext-01926  
  If the `pNext` chain includes a
  [VkRenderPassInputAttachmentAspectCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassInputAttachmentAspectCreateInfo.html)
  structure, the `subpass` member of each element of its
  `pAspectReferences` member **must** be less than `subpassCount`

- <a href="#VUID-VkRenderPassCreateInfo-pNext-01927"
  id="VUID-VkRenderPassCreateInfo-pNext-01927"></a>
  VUID-VkRenderPassCreateInfo-pNext-01927  
  If the `pNext` chain includes a
  [VkRenderPassInputAttachmentAspectCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassInputAttachmentAspectCreateInfo.html)
  structure, the `inputAttachmentIndex` member of each element of its
  `pAspectReferences` member **must** be less than the value of
  `inputAttachmentCount` in the element of `pSubpasses` identified by
  its `subpass` member

- <a href="#VUID-VkRenderPassCreateInfo-pNext-01963"
  id="VUID-VkRenderPassCreateInfo-pNext-01963"></a>
  VUID-VkRenderPassCreateInfo-pNext-01963  
  If the `pNext` chain includes a
  [VkRenderPassInputAttachmentAspectCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassInputAttachmentAspectCreateInfo.html)
  structure, for any element of the `pInputAttachments` member of any
  element of `pSubpasses` where the `attachment` member is not
  `VK_ATTACHMENT_UNUSED`, the `aspectMask` member of the corresponding
  element of
  [VkRenderPassInputAttachmentAspectCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassInputAttachmentAspectCreateInfo.html)::`pAspectReferences`
  **must** only include aspects that are present in images of the format
  specified by the element of `pAttachments` at `attachment`

- <a href="#VUID-VkRenderPassCreateInfo-pNext-01928"
  id="VUID-VkRenderPassCreateInfo-pNext-01928"></a>
  VUID-VkRenderPassCreateInfo-pNext-01928  
  If the `pNext` chain includes a
  [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)
  structure, and its `subpassCount` member is not zero, that member
  **must** be equal to the value of `subpassCount`

- <a href="#VUID-VkRenderPassCreateInfo-pNext-01929"
  id="VUID-VkRenderPassCreateInfo-pNext-01929"></a>
  VUID-VkRenderPassCreateInfo-pNext-01929  
  If the `pNext` chain includes a
  [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)
  structure, if its `dependencyCount` member is not zero, it **must** be
  equal to `dependencyCount`

- <a href="#VUID-VkRenderPassCreateInfo-pNext-01930"
  id="VUID-VkRenderPassCreateInfo-pNext-01930"></a>
  VUID-VkRenderPassCreateInfo-pNext-01930  
  If the `pNext` chain includes a
  [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)
  structure, for each non-zero element of `pViewOffsets`, the
  `srcSubpass` and `dstSubpass` members of `pDependencies` at the same
  index **must** not be equal

- <a href="#VUID-VkRenderPassCreateInfo-pNext-02512"
  id="VUID-VkRenderPassCreateInfo-pNext-02512"></a>
  VUID-VkRenderPassCreateInfo-pNext-02512  
  If the `pNext` chain includes a
  [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)
  structure, for any element of `pDependencies` with a `dependencyFlags`
  member that does not include `VK_DEPENDENCY_VIEW_LOCAL_BIT`, the
  corresponding element of the `pViewOffsets` member of that
  [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)
  instance **must** be `0`

- <a href="#VUID-VkRenderPassCreateInfo-pNext-02513"
  id="VUID-VkRenderPassCreateInfo-pNext-02513"></a>
  VUID-VkRenderPassCreateInfo-pNext-02513  
  If the `pNext` chain includes a
  [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)
  structure, elements of its `pViewMasks` member **must** either all be
  `0`, or all not be `0`

- <a href="#VUID-VkRenderPassCreateInfo-pNext-02514"
  id="VUID-VkRenderPassCreateInfo-pNext-02514"></a>
  VUID-VkRenderPassCreateInfo-pNext-02514  
  If the `pNext` chain includes a
  [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)
  structure, and each element of its `pViewMasks` member is `0`, the
  `dependencyFlags` member of each element of `pDependencies` **must**
  not include `VK_DEPENDENCY_VIEW_LOCAL_BIT`

- <a href="#VUID-VkRenderPassCreateInfo-pNext-02515"
  id="VUID-VkRenderPassCreateInfo-pNext-02515"></a>
  VUID-VkRenderPassCreateInfo-pNext-02515  
  If the `pNext` chain includes a
  [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)
  structure, and each element of its `pViewMasks` member is `0`, its
  `correlationMaskCount` member **must** be `0`

- <a href="#VUID-VkRenderPassCreateInfo-pDependencies-00837"
  id="VUID-VkRenderPassCreateInfo-pDependencies-00837"></a>
  VUID-VkRenderPassCreateInfo-pDependencies-00837  
  For any element of `pDependencies`, if the `srcSubpass` is not
  `VK_SUBPASS_EXTERNAL`, all stage flags included in the `srcStageMask`
  member of that dependency **must** be a pipeline stage supported by
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-types"
  target="_blank" rel="noopener">pipeline</a> identified by the
  `pipelineBindPoint` member of the source subpass

- <a href="#VUID-VkRenderPassCreateInfo-pDependencies-00838"
  id="VUID-VkRenderPassCreateInfo-pDependencies-00838"></a>
  VUID-VkRenderPassCreateInfo-pDependencies-00838  
  For any element of `pDependencies`, if the `dstSubpass` is not
  `VK_SUBPASS_EXTERNAL`, all stage flags included in the `dstStageMask`
  member of that dependency **must** be a pipeline stage supported by
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#synchronization-pipeline-stages-types"
  target="_blank" rel="noopener">pipeline</a> identified by the
  `pipelineBindPoint` member of the destination subpass

- <a href="#VUID-VkRenderPassCreateInfo-pDependencies-06866"
  id="VUID-VkRenderPassCreateInfo-pDependencies-06866"></a>
  VUID-VkRenderPassCreateInfo-pDependencies-06866  
  For any element of `pDependencies`, if its `srcSubpass` is not
  `VK_SUBPASS_EXTERNAL`, it **must** be less than `subpassCount`

- <a href="#VUID-VkRenderPassCreateInfo-pDependencies-06867"
  id="VUID-VkRenderPassCreateInfo-pDependencies-06867"></a>
  VUID-VkRenderPassCreateInfo-pDependencies-06867  
  For any element of `pDependencies`, if its `dstSubpass` is not
  `VK_SUBPASS_EXTERNAL`, it **must** be less than `subpassCount`

Valid Usage (Implicit)

- <a href="#VUID-VkRenderPassCreateInfo-sType-sType"
  id="VUID-VkRenderPassCreateInfo-sType-sType"></a>
  VUID-VkRenderPassCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO`

- <a href="#VUID-VkRenderPassCreateInfo-pNext-pNext"
  id="VUID-VkRenderPassCreateInfo-pNext-pNext"></a>
  VUID-VkRenderPassCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html),
  [VkRenderPassInputAttachmentAspectCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassInputAttachmentAspectCreateInfo.html),
  or
  [VkRenderPassMultiviewCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassMultiviewCreateInfo.html)

- <a href="#VUID-VkRenderPassCreateInfo-sType-unique"
  id="VUID-VkRenderPassCreateInfo-sType-unique"></a>
  VUID-VkRenderPassCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkRenderPassCreateInfo-flags-parameter"
  id="VUID-VkRenderPassCreateInfo-flags-parameter"></a>
  VUID-VkRenderPassCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkRenderPassCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateFlagBits.html) values

- <a href="#VUID-VkRenderPassCreateInfo-pAttachments-parameter"
  id="VUID-VkRenderPassCreateInfo-pAttachments-parameter"></a>
  VUID-VkRenderPassCreateInfo-pAttachments-parameter  
  If `attachmentCount` is not `0`, `pAttachments` **must** be a valid
  pointer to an array of `attachmentCount` valid
  [VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription.html) structures

- <a href="#VUID-VkRenderPassCreateInfo-pSubpasses-parameter"
  id="VUID-VkRenderPassCreateInfo-pSubpasses-parameter"></a>
  VUID-VkRenderPassCreateInfo-pSubpasses-parameter  
  `pSubpasses` **must** be a valid pointer to an array of `subpassCount`
  valid [VkSubpassDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription.html) structures

- <a href="#VUID-VkRenderPassCreateInfo-pDependencies-parameter"
  id="VUID-VkRenderPassCreateInfo-pDependencies-parameter"></a>
  VUID-VkRenderPassCreateInfo-pDependencies-parameter  
  If `dependencyCount` is not `0`, `pDependencies` **must** be a valid
  pointer to an array of `dependencyCount` valid
  [VkSubpassDependency](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDependency.html) structures

- <a href="#VUID-VkRenderPassCreateInfo-subpassCount-arraylength"
  id="VUID-VkRenderPassCreateInfo-subpassCount-arraylength"></a>
  VUID-VkRenderPassCreateInfo-subpassCount-arraylength  
  `subpassCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAttachmentDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentDescription.html),
[VkRenderPassCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSubpassDependency](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDependency.html),
[VkSubpassDescription](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription.html),
[vkCreateRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
