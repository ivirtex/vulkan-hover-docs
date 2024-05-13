# VK_AMD_shader_info(3) Manual Page

## Name

VK_AMD_shader_info - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

43

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">Developer tools</a>

## <a href="#_contact" class="anchor"></a>Contact

- Jaakko Konttinen <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_shader_info%5D%20@jaakkoamd%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_shader_info%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jaakkoamd</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2017-10-09

**IP Status**  
No known IP claims.

**Contributors**  
- Jaakko Konttinen, AMD

## <a href="#_description" class="anchor"></a>Description

This extension adds a way to query certain information about a compiled
shader which is part of a pipeline. This information may include shader
disassembly, shader binary and various statistics about a shader’s
resource usage.

While this extension provides a mechanism for extracting this
information, the details regarding the contents or format of this
information are not specified by this extension and may be provided by
the vendor externally.

Furthermore, all information types are optionally supported, and users
should not assume every implementation supports querying every type of
information.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetShaderInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetShaderInfoAMD.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkShaderResourceUsageAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderResourceUsageAMD.html)

- [VkShaderStatisticsInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderStatisticsInfoAMD.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkShaderInfoTypeAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderInfoTypeAMD.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_AMD_SHADER_INFO_EXTENSION_NAME`

- `VK_AMD_SHADER_INFO_SPEC_VERSION`

## <a href="#_examples" class="anchor"></a>Examples

This example extracts the register usage of a fragment shader within a
particular graphics pipeline:

``` c
extern VkDevice device;
extern VkPipeline gfxPipeline;

PFN_vkGetShaderInfoAMD pfnGetShaderInfoAMD = (PFN_vkGetShaderInfoAMD)vkGetDeviceProcAddr(
    device, "vkGetShaderInfoAMD");

VkShaderStatisticsInfoAMD statistics = {};

size_t dataSize = sizeof(statistics);

if (pfnGetShaderInfoAMD(device,
    gfxPipeline,
    VK_SHADER_STAGE_FRAGMENT_BIT,
    VK_SHADER_INFO_TYPE_STATISTICS_AMD,
    &dataSize,
    &statistics) == VK_SUCCESS)
{
    printf("VGPR usage: %d\n", statistics.resourceUsage.numUsedVgprs);
    printf("SGPR usage: %d\n", statistics.resourceUsage.numUsedSgprs);
}
```

The following example continues the previous example by subsequently
attempting to query and print shader disassembly about the fragment
shader:

``` c
// Query disassembly size (if available)
if (pfnGetShaderInfoAMD(device,
    gfxPipeline,
    VK_SHADER_STAGE_FRAGMENT_BIT,
    VK_SHADER_INFO_TYPE_DISASSEMBLY_AMD,
    &dataSize,
    nullptr) == VK_SUCCESS)
{
    printf("Fragment shader disassembly:\n");

    void* disassembly = malloc(dataSize);

    // Query disassembly and print
    if (pfnGetShaderInfoAMD(device,
        gfxPipeline,
        VK_SHADER_STAGE_FRAGMENT_BIT,
        VK_SHADER_INFO_TYPE_DISASSEMBLY_AMD,
        &dataSize,
        disassembly) == VK_SUCCESS)
    {
        printf((char*)disassembly);
    }

    free(disassembly);
}
```

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2017-10-09 (Jaakko Konttinen)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_AMD_shader_info"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
